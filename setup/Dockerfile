FROM homebrew/brew

RUN mkdir keenan && cd keenan
WORKDIR /keenan/

COPY ./ /keenan/docs

# install swagger
RUN brew update-reset
RUN brew tap go-swagger/go-swagger
RUN brew install go-swagger

RUN cd /keenan/docs && swagger validate ./swagger.yml

CMD [ "swagger", "serve", "/keenan/docs/swagger.yml", "--port=8070", "--no-open"]