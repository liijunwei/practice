class Notifier
  def do_notice(event_type, target)
    event_type.notice(target)
  end
end

class Alert
  def notice(target)
    puts "notice alert via phone to #{target}"
  end
end

class AlertOk
  def notice(target)
    puts "notice alert_ok via im to #{target}"
  end
end

notifier = Notifier.new

notifier.do_notice(Alert.new, "176yyyyxxxx")
notifier.do_notice(AlertOk.new, "532388888")
