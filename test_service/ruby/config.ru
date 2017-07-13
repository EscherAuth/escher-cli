require 'rack/app'

class App < Rack::App

  desc 'health check endpoint'
  get '/' do
    'OK'
  end

  get '/hello' do
    puts payload
    'response body'
  end

end

run App
