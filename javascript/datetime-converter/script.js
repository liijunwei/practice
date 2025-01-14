
// https://momentjs.com/docs/
setInterval(() => {
  let now = moment()
  document.querySelector(".js-date1").innerHTML = now.format();
  document.querySelector(".js-date2").innerHTML = now.utc().format();
  document.querySelector(".js-date3").innerHTML = now.utc().toISOString();
}, 1000);

// console.log(moment(new Date()))
// console.log(moment().utcOffset(-5))
// console.log(moment().utcOffset(-5))
// console.log(now.toISOString())
