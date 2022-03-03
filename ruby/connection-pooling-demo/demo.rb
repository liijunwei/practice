require 'mysql2'
require 'connection_pool'
require 'pry'

database = ENV['database']

# https://github.com/mperham/connection_pool/
pool = ConnectionPool.new(size: 5, timeout: 5) {
  # https://github.com/brianmario/mysql2/#connection-options
  Mysql2::Client.new(
    host:         "localhost",
    username:     ENV['username'],
    password:     ENV['password'],
    database:     database,
    encoding:     'utf8',
    init_command: "use #{database}"
  )
}

# result = pool.with {|conn| conn.query("SELECT sleep(5);")}
# binding.pry

threads = []

num = 5
# num = 10

num.times do |i|
  threads << Thread.new {pool.with {|conn| conn.query("SELECT sleep(5);")}}
end

binding.pry
# puts result.to_a


