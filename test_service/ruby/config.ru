require 'rack/app'
require 'yaml'

class App < Rack::App

  desc 'health check endpoint'
  get '/' do
    STDOUT.puts(YAML.dump(request))
    'OK'
  end

  get '/hello' do
    puts payload
    'response body'
  end

end

run App
