+ Extraneous code adds costs without providing benefits, and at this point, it’s quite reasonable to challenge the need for song.

+ Does song serve a purpose independent of verses, or is it redundant and thus a candidate for deletion?
+ Answering this question requires thinking about the problem from the message sender’s point of view.

+ The song method imposes a single dependency; to use it, you need only know its name.

song方法是API

+ Using the verses method to request the entire song, however, requires significantly more knowledge. The sender must know:
    + the name of the verses method
    + that the method requires two arguments
    + that the first argument is the verse on which to start that the second argument is the verse on which to end
    + that the song starts on verse 99 that the song ends on verse 0

+ This is a lot of knowledge.
+ There are many ways in which the verses method could change that would break senders of this message.


