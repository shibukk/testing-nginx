FROM ruby:latest

RUN mkdir -p /app
WORKDIR /app
COPY ./ruby .

RUN gem update --system --no-ri --no-rdoc && bundle install --jobs=4

CMD ["ruby", "main.rb", "-p", "8080", "-o", "0.0.0.0"]
