package mail

import (
	"net/smtp"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type smtpServer struct {
	host string
	port string
}

func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

func SendOTP(toRollNo string, otp string) (err error) {
	log.Debug("<:ohyeah:897524824769974283> <:1red:897524876737392650> <:point2:895604250573410345> Sending OTP ", otp, " to ", toRollNo, " <:coin:892312139161821195> ")
	from := os.Getenv("MAIL_ID")
	password := os.Getenv("MAIL_PASSWORD")
	to := []string{
		toRollNo + "@iitk.ac.in",
	}
	smtpServer := smtpServer{host: viper.GetString("MAIL.HOST"), port: viper.GetString("MAIL.PORT")}

	message := []byte("Your OTP is " + otp)

	auth := smtp.PlainAuth("", from, password, smtpServer.host)

	err = smtp.SendMail(smtpServer.Address(), auth, from, to, message)

	if err != nil {
		log.Error("error sending mail: ", err)
		return err
	}
	log.Info("Mail sent to ", toRollNo)
	return nil
}
