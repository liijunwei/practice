You want to interact with a complex API of an external system or resource in a simplified way.

*Introduce a Gateway that encapsulates access to an external system or resource.*

+ The majority of Rails applications use ActiveRecord as a Gateway to a relational database.

+ **APIs are naturally going to be somewhat complicated because they are designed to be flexible and reusable for various consumers.**

```ruby
class Gateway
  attr_accessor :subject, :attributes, :to

  def self.save
    gateway = self.new
    yield gateway
    gateway.execute
  end

  def execute
    request = Net::HTTP::Post.new(url.path)

    attribute_hash = attributes.inject({}) do |result, attribute|
      result[attribute.to_s] = subject.send attribute
      result
    end

    request.set_form_data(attribute_hash)
    Net::HTTP.new(url.host, url.port).start {|http| http.request(request) }
  end

  def url
    URI.parse(to)
  end
end

class Person
  attr_accessor :first_name, :last_name, :ssn

  def save
    Gateway.save do |persist|
      persist.subject = self
      persist.attributes = [:first_name, :last_name, :ssn]
      persist.to = 'http://www.example.com/person'
    end
  end
end

class PostGateway < Gateway
  # ...
end

class GetGateway < Gateway
  # ...
end
```
