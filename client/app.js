const btn = document.getElementById('numberBtn')
const input = document.getElementById('numberInput')
const output = document.getElementById('output')
const spinner = document.getElementById("spinner");

btn.addEventListener('click', () => {
  let data = input.value
  
  fetch('http://localhost:8080/pi/' + data)
  .then((res) => res.json())
  .then((data) => {
    
    clean()
    let index = data['index']
    let msg =  data['text'] + ' ...'
    if(index != 0) {
      msg = '...' + msg
    }

    output.innerHTML = "<h2>your number was found at index " + index + "</h2><br>"
    output.innerHTML += clearStr(msg)
  })
  .catch(() => alert('something went wrong, try again'))
})

input.addEventListener('keyup', (event) => {
  if(event.keyCode === 13) {
    event.preventDefault();
    btn.click()
  }
})

function clean() {
  input.value = ''
  output.innerHTML = ''
}

function clearStr(msg) {
  let tmp = msg.replace('s', ' <span class=\'num\'>')
  let final = tmp.replace('e', '</span>')
  return final
}