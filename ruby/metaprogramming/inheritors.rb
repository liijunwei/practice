# page 244

$INHERITORS = []

class C
  def self.inherited(subclass)
   $INHERITORS << subclass
  end
end

class D < C ;end
class E < C ;end
class F < C ;end
class G < C ;end

p $INHERITORS
