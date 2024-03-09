export function createPVPCavElement(_container, _dpr, _isRandomIndex, _index){
  if(_isRandomIndex){
    let t = new Date().getTime().toString()
    _index = parseInt(t.substring(t.length - 6, t.length)) + _index
  }
  const canvas = document.createElement('canvas')
  canvas.style.width = "380px"
  canvas.style.height = "230px"
  canvas.className  = 'battlecav skill'
  canvas.width = 380 * _dpr;
  canvas.height = 230 * _dpr;
  canvas.style.zIndex = _index
  _container.appendChild(canvas)
  return canvas
}

export function createPVPCav(_cavElement, _dpr){
  let cav = _cavElement.getContext('2d');
  cav.scale(_dpr, _dpr)
  return cav
}
