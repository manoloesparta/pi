// Elements
const btn = document.getElementById('numberBtn')
const input = document.getElementById('numberInput')
const output = document.getElementById('output')
const loading = document.getElementById('lds-ring')

// Functions
const clean = () => {
  output.innerHTML = ''
  input.value = input.value.trim().replace(/\D+/g, '')
}

const randomDecimal = () => Math.floor(Math.random() * 10)
const random256 = () => crypto.getRandomValues(new Uint8Array(256)).join('').substring(0,256).split('').join(' ')

// Event listener
btn.addEventListener('click', async () => {

  clean()
  if(input.value.trim() == '') return

  try {

    const inputVal = input.value
    loading.classList.remove('hidden')

    const response = await fetch(`https://piapi.manoloesparta.com/pi/${inputVal}`)
    const json = await response.json()

    let { index, text } = json

    if (index != 0) text = '...' + text
    output.innerHTML = `<h2>your number was found at index ${index}</h2><br>`
    output.innerHTML += text.replace('s', ' <span class=\'num\'>').replace('e', '</span>')
    
  } catch {
    alert('something went wrong, try again')
  }

  loading.classList.add('hidden')
})

// Use key enter to search
input.addEventListener('keyup', (event) => {
  if(event.keyCode === 13) {
    event.preventDefault();
    btn.click()
  }
})
