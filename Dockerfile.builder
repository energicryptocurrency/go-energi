# Ubuntu 18.04 is our base image for building
FROM ubuntu:18.04

# set up timezone
ENV TZ=GMT
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# update software
RUN apt -y --fix-missing update
RUN apt -y full-upgrade
RUN apt -y autoremove
RUN apt -y clean

# install docker
RUN apt -y update
RUN apt -y install curl gnupg lsb-release software-properties-common
RUN curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
RUN add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
RUN apt -y install docker-ce docker-ce-cli containerd.io

# install development tools
RUN apt -y install git vim htop apg jq direnv build-essential wget awscli sudo

# golang variables
ARG golang_version="1.15.8"
ARG golang_hostarch="linux-amd64"
ARG golang_filename="go${golang_version}.${golang_hostarch}.tar.gz"
ARG golang_url="https://golang.org/dl/${golang_filename}"
ARG golang_sha256="d3379c32a90fdf9382166f8f48034c459a8cc433730bc9476d39d9082c94583b"

# install golang
RUN wget -nv ${golang_url}
RUN echo "${golang_sha256} ${golang_filename}" > "${golang_filename}.sha256"
RUN sha256sum -c ${golang_filename}.sha256
RUN tar -C /usr/local -xzf ${golang_filename}
RUN rm -rf ${golang_filename}*
ENV PATH="${PATH}:/usr/local/go/bin"
ENV GOROOT="/usr/local/go"

# install go-junit-report
RUN go get -u -v github.com/RyanLucchese/go-junit-report
ENV PATH="${PATH}:/root/go/bin"

# nodejs variables
ARG nodejs_version="12.22.1"
ARG nodejs_hostarch="linux-x64"
ARG nodejs_spec="node-v${nodejs_version}-${nodejs_hostarch}"
ARG nodejs_filename="${nodejs_spec}.tar.gz"
ARG nodejs_url="https://nodejs.org/download/release/v12.22.1/${nodejs_filename}"
ARG nodejs_sha256="d315c5dea4d96658164cdb257bd8dbb5e44bdd2a7c1d747841f06515f23a0042"

# install nodejs
RUN wget -nv ${nodejs_url}
RUN echo "${nodejs_sha256} ${nodejs_filename}" > "${nodejs_filename}.sha256"
RUN sha256sum -c ${nodejs_filename}.sha256
RUN tar -C /usr/local -xzf ${nodejs_filename}
RUN chown -R "root:root" "/usr/local/${nodejs_spec}"
RUN rm -rf ${nodejs_filename}*
ENV PATH="${PATH}:/usr/local/${nodejs_spec}/bin"

# install node packages
RUN npm -g config set user root
RUN npm install -g yarn ganache-cli truffle

# /builds/energi/tech/gen3/energi3
RUN mkdir -p "/builds/energi/tech/gen3"
WORKDIR "/builds/energi/tech/gen3"
WORKDIR "/builds/energi/tech/gen3/energi3"
ADD Makefile.release Makefile.release
ADD package.json package.json
RUN npm install
RUN make -f Makefile.release release-tools
ENV GOPATH="/builds/energi/tech/gen3"
ENV GOBIN="/builds/energi/tech/gen3/energi3/build/bin"
ENV GO111MODULE="on"
ENV GOFLAGS="-mod=vendor -v"

# do a build at the end to ensure we have everything
RUN make all
# TODO: make check is known to fail now due to issues in tests and the linter
#RUN make check
