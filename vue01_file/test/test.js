
const { level } = 1
let pathlist = [{'aa': 1, 'bb': 2}, {'aa': 3, 'bb': 4}]

let pathLabels = ['asd', 'asdasd']
pathLabels.splice(0, 1, '/')
let aa = '/' + pathLabels.join('/')
const nodes = Array.from(pathlist)
  .map(item => ({
    value: item.bb,
    label: item.aa,
    leaf: level >= 2
  }))
const nodes1 = Array.from(pathlist)
  .map(item => (item.aa))
console.log(nodes1)
