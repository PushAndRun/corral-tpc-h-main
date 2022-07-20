module github.com/tawalaya/corral_plus_tpch

go 1.16

require (
	github.com/ISE-SMILE/corral v0.2.2
	github.com/dnlo/struct2csv v0.0.0-20190928115744-2f584471b24e // indirect
	github.com/go-git/go-git/v5 v5.4.2
	github.com/google/martian/v3 v3.2.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/viper v1.12.0
)

replace github.com/ISE-SMILE/corral v0.2.2 => ../corral
