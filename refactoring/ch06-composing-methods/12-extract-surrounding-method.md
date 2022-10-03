# Extract Surrounding Method

You have two methods that contain nearly identical code. The variance is in the middle of the method.

**Extract the duplication into a method that accepts a block and yields back to the caller to execute the unique code.** (mind blowing)

```ruby
# before
def charge(amount, credit_card_number)
  begin
    connection = CreditCardServer.connect(...)
    connection.send(amount, credit_card_number)
  rescue IOError => e
    Logger.log "Could not submit order #{@order_number} to the server: #{e}"
    return nil
  ensure
    connection.close
  end
end

# after
def charge(amount, credit_card_number)
  connect do |connection|
    connection.send(amount, credit_card_number)
  end
end

def connect
  begin
    connection = CreditCardServer.connect(...)
    yield connection
  rescue IOError => e
    Logger.log("Could not submit order...")
    return nil
  ensure
    connection.close
  end
end
```

+ As well as removing duplication, this refactoring can be used to hide away infrastructure code (for example, code for iterating over a collection or connecting to an external service), so that the business logic becomes more prominent.

如果两段代码只有中间部分是不同的, 可以用 ruby block 来抽象(one of the block usages)
