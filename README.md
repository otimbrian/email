Instructions.

1. Start by setting up your IMAP and turning it on in GMAIL.
follow instructions <a href="https://support.google.com/a/answer/9003945#imap_gmail&zippy=%2Cstep-turn-on-imap-in-gmail
">Here</a>.

2. Create the pplication password to be used in sending the Emails.
follow instructions <a hrref="https://support.google.com/mail/answer/185833">Here</a>

3. Initialize gomail.
    m := gomail.NewMessage()

4. Add the sender email and the receiving email(s) in the header
    ReciepientAndUser(m, "sender@email.com", "receiver@email.com")

    if they are many, 

    ReciepientAndUser(m, "sender@email.com", "receiver1@email.com", "receiver2@email.com")

5. Create messsage with subject and body.
    CreateMessageBody(m, "Subject", "Body")

6. Send
    message.SendEmail(m, "sender@email.com", password)
    Where password is the application password created in step 2.
