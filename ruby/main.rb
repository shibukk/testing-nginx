require 'sinatra'

get '/' do
  response.headers["X-APP-USER-ID"] = "12345"
  'Home'
end

get '/login' do
  'Login'
end