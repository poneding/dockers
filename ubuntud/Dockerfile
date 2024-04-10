FROM ubuntu:22.04

LABEL author="poneding@gmail.com"

RUN apt-get update && apt-get install -y openssh-server && mkdir /var/run/sshd
RUN useradd ubuntu && echo "ubuntu:ubuntu" | chpasswd && \
    echo "ubuntu ALL=(ALL)" >> /etc/sudoers && mkdir /home/ubuntu

EXPOSE 22
ENTRYPOINT /usr/sbin/sshd -D