require 'mysql2'
require 'pry'

# https://github.com/brianmario/mysql2/#connection-options
client = Mysql2::Client.new(
  :host => "localhost",
  :username => ENV['username'],
  :password => ENV['password'],
  :database => ENV['database'],
  :encoding => 'utf8',
  :init_command => "use #{ENV['database']}"
)

result = client.query("SELECT * FROM reasons limit 10")

binding.pry
puts result.to_a


