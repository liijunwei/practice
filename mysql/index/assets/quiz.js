/* Reusable self-check quiz component.
   Markup contract:
     <div class="quiz">
       <div class="question" data-correct="<value>">
         <p class="q">Question text</p>
         <button class="option" data-value="<value>">Option text</button>
         ...
         <p class="expl">Explanation shown after answering.</p>
       </div>
     </div>
   Each .option's data-value is compared to the question's data-correct.
   First click locks the question and reveals correct/wrong + explanation.
*/
(function () {
  function init() {
    document.querySelectorAll('.quiz').forEach(function (quiz) {
      quiz.querySelectorAll('.question').forEach(function (q) {
        var correct = q.getAttribute('data-correct');
        var expl = q.querySelector('.expl');
        q.querySelectorAll('.option').forEach(function (opt) {
          opt.addEventListener('click', function () {
            if (q.classList.contains('answered')) return;
            q.classList.add('answered');
            var chosen = opt.getAttribute('data-value');
            q.querySelectorAll('.option').forEach(function (o) {
              o.disabled = true;
              if (o.getAttribute('data-value') === correct) o.classList.add('correct');
            });
            if (chosen !== correct) opt.classList.add('wrong');
            if (expl) expl.classList.add('show');
          });
        });
      });
    });
  }
  if (document.readyState === 'loading') {
    document.addEventListener('DOMContentLoaded', init);
  } else {
    init();
  }
})();
