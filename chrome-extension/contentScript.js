// TODO: allow customizations
const defaultSeperator = "*****new case*****\n"

// DOM -> test case content
function parseCaseField(dom) {
  let text = dom.innerHTML.trim().split(/<br>/).join("\n")

  // Decode html entities
  const div = document.createElement("div")
  div.innerHTML = text
  text = div.textContent

  // For Windows
  return text.replace(/\n/g, "\r\n")
}

// Giving a selector & seperator produce the case group content for .in or .out
function genCaseGroup(selector, seperator) {
  const doms = Array.from(document.querySelectorAll(selector))
  return doms.map(d => parseCaseField(d)).join(seperator)
}

function download(content, name) {
  const a = document.createElement("a")
  a.setAttribute("href", "data:text/plain," + encodeURIComponent(content))
  a.setAttribute("download", name)

  document.body.appendChild(a)
  a.click()
}

const title = document.querySelector(".header .title").textContent[0]

const button = document.createElement("button")
document.querySelector(".header").appendChild(button)
button.innerHTML = "Download Cases"
button.addEventListener("click", () => {
  download(genCaseGroup(".input pre", defaultSeperator), `${title}.in`)
  download(genCaseGroup(".output pre", defaultSeperator), `${title}.out`)
})