const fwdays = () => {
  const fwstart = new Date('2022-02-24')
  const now = new Date()

  return Math.ceil((now - fwstart) / (1000 * 60 * 60 * 24))
}

const owdays = () => {
  const owstart = new Date('2014-02-26')
  const now = new Date()

  return Math.ceil((now - owstart) / (1000 * 60 * 60 * 24))
}

(function() {
  document.getElementById('uwdays').innerHTML = `Already ${fwdays()} days of Russian fullscale war on Ukraine. 
  ${owdays()} days since Russian occupational war on Ukraine's Crimea and Donbas. <a href="https://standforukraine.com">Stand for Ukraine! Help Ukrainians to defeat Russia!</a> `
})()
