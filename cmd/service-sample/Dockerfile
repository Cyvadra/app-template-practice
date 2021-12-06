FROM jrei/systemd-centos:7

COPY AppPractice.viper.yaml /etc/AppPractice/AppPractice.viper.yaml
COPY app-practice /usr/local/bin/app-practice

CMD ["app-practice", "run"]
