FROM buildpack-deps:buster

RUN apt-get update && apt-get install -y luajit lua5.1 lua5.1-dev libxss1 libappindicator1 libindicator7

RUN wget https://luarocks.org/releases/luarocks-3.3.1.tar.gz \
    && tar zxpf luarocks-3.3.1.tar.gz \
    && cd luarocks-3.3.1 \
    && ./configure && make && make install

RUN luarocks install --server=https://luarocks.org/dev luaffi

RUN wget https://dl.google.com/linux/direct/google-chrome-stable_current_amd64.deb && \
    apt install -fy ./google-chrome*.deb

COPY . /app

WORKDIR /app
CMD ["luajit", "main.lua"]