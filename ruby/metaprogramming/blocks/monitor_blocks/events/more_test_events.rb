# what if we want to share variable between events?

def monthly_sales
  110 # read productoin data from database
end

target_sales = 100

event "monthly sales are suspicioutly high" do
  monthly_sales > target_sales
end

event "monthly sales are absymally low" do
  monthly_sales < target_sales
end

