FROM ruby:4.0.0

RUN apt-get update -qq && apt-get install -y nodejs postgresql-client libpq-dev

WORKDIR /app

COPY Gemfile Gemfile.lock ./

RUN bundle install

COPY . .

EXPOSE 3001

CMD ["rails", "server", "-b", "0.0.0.0", "-p", "3001"]
