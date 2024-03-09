export function FallItem(_x, _y, _itemId, _cavCtx, _element, _dpr){
  this.xSpd = Math.floor(Math.random() * 21) + (-10);
  this.ySpd = Math.floor(Math.random() * 10) + (-30);
  this.yEnd = Math.floor(Math.random() * 30) + 190;
  this.x = _x
  this.y = _y
  this.dpr = _dpr
  this.fadexEnd = _element.width / _dpr - 25
  this.fadeyEnd = 0
  this.fadetime = 20;
  this.fxSpd = 0
  this.fadexStart = 0
  this.fySpd = 0
  this.fadeyStart = 0
  this.scale = 1
  this.itemId = _itemId
  this.canvas = _cavCtx
  this.element = _element
  this.inFade = false

  this.draw = function(){
    let img = new Image()
    img.src = require('@/assets/items/'+ this.itemId + (parseInt(this.itemId) === 10 ? 'b' : '') +'.png')
    img.onload = (() => {
      this.canvas.clearRect(0, 0, this.element.width, this.element.height);
      this.canvas.shadowOffsetY = 2
      this.canvas.shadowBlur = 2
      this.canvas.shadowColor = 'rgb(0,0,0,0.8)'
      this.canvas.drawImage(img, this.x, this.y, 20 * this.scale, 20 * this.scale)
    })
  }

  this.fall = () =>{
    this.draw()
    let bool = this.updateFall()
    if(bool){
      requestAnimationFrame(this.fall)
    }
    console.log("do fall====")
  }

  this.fade = () =>{
    this.draw()
    let bool = this.updateFade()
    if(bool){
      requestAnimationFrame(this.fade)
    }
    console.log("do fade====")
  }

  this.updateFall = function(){
    if(this.inFade){
      return false
    }
    if(this.y < this.yEnd){
      this.ySpd += 5
      this.x += this.xSpd
      this.y += this.ySpd
    }else{
      this.x += 0
      this.y += 0
      return false
    }
    this.fadexStart = this.x
    this.fadeyStart = this.y
    return true
  }

  this.updateFade = function(){
    this.inFade = true
    this.fxSpd = (this.fadexEnd - this.fadexStart) / this.fadetime
    this.fySpd = (this.fadeyEnd - this.fadeyStart) / this.fadetime
    this.x += this.fxSpd
    this.y += this.fySpd
    this.scale = this.scale <= 0.8 ? 0.8 : (this.scale - 0.05)
    if(this.y <= this.fadeyEnd){
      this.element.remove()
      return false
    }
    return true
  }

}
