package email

import "testing"

func TestSendEmail(t *testing.T) {
	type args struct {
		subject string
		from    string
		to      []string
		cc      []string
		content string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"linx test email ",
			args{
				subject: "this is test",
				from:    "1824517828@qq.com",
				to:      []string{"xionglin@tuochexia.cn"},
				cc:      []string{"782626063@qq.com"},
				content: "你是不是狗",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendEmail(tt.args.subject, tt.args.from, tt.args.to, tt.args.cc, tt.args.content)
		})
	}
}

func TestSendEmailHtml(t *testing.T) {
	type args struct {
		subject string
		from    string
		to      []string
		cc      []string
		htmlStr string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"linx test email ",
			args{
				subject: "this is test",
				from:    "1824517828@qq.com",
				to:      []string{"xionglin@tuochexia.cn"},
				cc:      []string{"782626063@qq.com"},
				htmlStr: "<h1>你是不是狗</h1>",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendEmailHtml(tt.args.subject, tt.args.from, tt.args.to, tt.args.cc, tt.args.htmlStr)
		})
	}
}

func Test_sendHtml(t *testing.T) {
	type args struct {
		subject string
		from    string
		to      []string
		cc      []string
		bcc     []string
		htmlStr string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sendHtml(tt.args.subject, tt.args.from, tt.args.to, tt.args.cc, tt.args.bcc, tt.args.htmlStr)
		})
	}
}

func Test_sendText(t *testing.T) {
	type args struct {
		subject string
		from    string
		to      []string
		cc      []string
		bcc     []string
		content string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sendText(tt.args.subject, tt.args.from, tt.args.to, tt.args.cc, tt.args.bcc, tt.args.content)
		})
	}
}
