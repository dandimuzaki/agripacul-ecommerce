package template

import (
	"fmt"
	"time"
)

func RequestResetPassword(name, otp, appName string) string {
return fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
  <title>Password Reset</title>
</head>
<body style="margin:0; padding:0; background-color:#f4f6f8; font-family: Arial, Helvetica, sans-serif;">
  <table width="100%%" cellpadding="0" cellspacing="0">
    <tr>
      <td align="center" style="padding: 24px;">
        <table width="600" cellpadding="0" cellspacing="0" style="background-color:#ffffff; border-radius:8px; overflow:hidden;">
          
          <!-- Header -->
          <tr>
            <td style="padding: 24px; text-align:center; background-color:#2563eb; color:#ffffff;">
              <h2 style="margin:0;">Reset Your Password</h2>
            </td>
          </tr>

          <!-- Body -->
          <tr>
            <td style="padding: 32px; color:#333333;">
              <p style="margin-top:0;">Hello <strong>%s</strong>,</p>

              <p>
                We received a request to reset the password for your account.
                Please use the One-Time Password (OTP) below to continue.
              </p>

              <div style="margin: 24px 0; text-align:center;">
                <span style="
                  display:inline-block;
                  padding: 14px 24px;
                  font-size: 24px;
                  letter-spacing: 4px;
                  font-weight: bold;
                  background-color:#f3f4f6;
                  border-radius:6px;
                ">
                  %s
                </span>
              </div>

              <p>
                This OTP will expire in <strong>10 minutes</strong>.
                If you did not request a password reset, please ignore this email.
              </p>

              <p style="color:#6b7280; font-size:14px;">
                For your security, do not share this OTP with anyone.
              </p>

              <p style="margin-bottom:0;">
                Regards,<br>
                <strong>%s Team</strong>
              </p>
            </td>
          </tr>

          <!-- Footer -->
          <tr>
            <td style="padding: 16px; text-align:center; background-color:#f9fafb; color:#9ca3af; font-size:12px;">
              © %d %s. All rights reserved.
            </td>
          </tr>

        </table>
      </td>
    </tr>
  </table>
</body>
</html>
`,
	name,
	otp,
	appName,
	time.Now().Year(),
	appName,
)
}