function getTime() {
  let now = new Date();
  // current date
  // adjust 0 before single digit date
  let date = ("0" + now.getDate()).slice(-2);
  // current month
  let month = ("0" + (now.getMonth() + 1)).slice(-2);
  // current year
  let year = now.getFullYear();
  // current hours
  let hours = ("0" + now.getHours()).slice(-2);
  // current minutes
  let minutes = ("0" + now.getMinutes()).slice(-2);
  // current seconds
  let seconds = ("0" + now.getSeconds()).slice(-2);

  return "["+year+"-"+month+"-"+date+" "+hours+":"+minutes+":"+seconds+"]";
}

module.exports = { getTime }
