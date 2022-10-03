`Remove the default value` when a parameter has a default value, but the method is never called without the parameter.

+ But sometimes, as code evolves over time, fewer and fewer callers require the default value, until finally the default value is unused.

+ **Unused flexibility in software is a bad thing.** Maintenance of this flexibility takes time, allows opportunities for bugs, and makes refactoring more difficult.
+ Unused default parameters should be removed.
