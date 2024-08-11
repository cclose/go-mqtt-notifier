data "aws_route53_zone" "main" {
    name = var.domain_name
}

resource "aws_route53_zone" "mail_subdomain" {
  name = var.subdomain_name
}

resource "aws_route53_record" "mail-ns" {
  zone_id = data.aws_route53_zone.main.zone_id
  name    = var.subdomain_name
  type    = "NS"
  ttl     = "30"
  records = aws_route53_zone.mail_subdomain.name_servers
}

resource "aws_ses_domain_identity" "mail_subdomain" {
  domain = var.subdomain_name
}

resource "aws_ses_domain_dkim" "mail_subdomain_dkim" {
  domain = aws_ses_domain_identity.mail_subdomain.domain
}

resource "aws_route53_record" "mail_verification" {
  zone_id = aws_route53_zone.mail_subdomain.zone_id
  name    = "_amazonses.${var.subdomain_name}"
  type    = "TXT"
  ttl     = "300"
  records = [aws_ses_domain_identity.mail_subdomain.verification_token]
}

resource "aws_route53_record" "mail_dkim" {
  count   = 3
  zone_id = aws_route53_zone.mail_subdomain.zone_id
  name    = "${aws_ses_domain_dkim.mail_subdomain_dkim.dkim_tokens[count.index]}._domainkey"
  type    = "CNAME"
  ttl     = "300"
  records = ["${aws_ses_domain_dkim.mail_subdomain_dkim.dkim_tokens[count.index]}.dkim.amazonses.com"]
}

resource "aws_route53_record" "mail_mx" {
  zone_id = aws_route53_zone.mail_subdomain.zone_id
  name    = var.subdomain_name
  type    = "MX"
  ttl     = "300"
  records = ["10 inbound-smtp.${data.aws_region.current.name}.amazonaws.com"]  # Adjust as per your SES region
}

resource "aws_route53_record" "dmarc" {
  zone_id = aws_route53_zone.mail_subdomain.zone_id
  name    = "_dmarc.${var.subdomain_name}"
  type    = "TXT"
  ttl     = "300"
  records = ["v=DMARC1; p=none;"]
}

resource "aws_ses_domain_identity_verification" "verify_mail_subdomain" {
  domain = aws_ses_domain_identity.mail_subdomain.domain

  depends_on = [
    aws_route53_record.mail_verification
  ]
}

resource "aws_ses_domain_mail_from" "mail_subdomain" {
  domain           = aws_ses_domain_identity.mail_subdomain.domain
  mail_from_domain = "bounce.${aws_ses_domain_identity.mail_subdomain.domain}"
}

resource "aws_route53_record" "mail_from_mx" {
  zone_id = aws_route53_zone.mail_subdomain.zone_id
  name    = aws_ses_domain_mail_from.mail_subdomain.mail_from_domain
  type    = "MX"
  ttl     = "300"
  records = ["10 feedback-smtp.${data.aws_region.current.name}.amazonses.com"]
}

resource "aws_route53_record" "mail_from_spf" {
  zone_id = aws_route53_zone.mail_subdomain.zone_id
  name    = aws_ses_domain_mail_from.mail_subdomain.mail_from_domain
  type    = "TXT"
  ttl     = "300"
  records = ["v=spf1 include:amazonses.com ~all"]
}
