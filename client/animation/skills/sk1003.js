import { createPVPCavElement, createPVPCav } from '@/animation/cavUtil'

export function Sk1003(_container, _dpr){
  this.skillStatus = 0 // 0 init  1 doing  99 finish
  this.dpr = _dpr
  this.element = createPVPCavElement(_container, _dpr, true, 0)
  this.cav = createPVPCav(this.element, _dpr)
  this.hitelement = createPVPCavElement(_container, _dpr, true, 0)
  this.hitcav = createPVPCav(this.hitelement, _dpr)

  this.element2 = createPVPCavElement(_container, _dpr, true, 0)
  this.cav2 = createPVPCav(this.element2, _dpr)
  this.hitelement2 = createPVPCavElement(_container, _dpr, true, 0)
  this.hitcav2 = createPVPCav(this.hitelement2, _dpr)

  this.do = function(_atker, _beAtker){
    this.skillStatus = 1
    let xOffset = -150
    let count = 0
    if(_beAtker.type == 1){
      // 1表示我被攻击
      this.cav.scale(-1, 1);
      this.cav.translate(-this.element.width / this.dpr, 0);
      xOffset = 30
    }
    let diffTime = 0
    let lastDate = Date.now()
    let idx = 0
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        img.src = require('@/assets/frameimg/skills/1003/cast_' + idx + '.png')
        img.onload = (() => {
          that.cav.clearRect(0, 0, that.element.width, that.element.height)
          that.cav.drawImage(img, _beAtker.xStand + xOffset, _beAtker.yStand + 7, 143, 80)
        })
        idx++
        xOffset = xOffset + 8
        if(count === 0 && idx === 12){
          that.doSecond(_atker, _beAtker)
        }
        if(idx === 15){
          count++
          if(count === 1){
            that.hit(_beAtker)
          }
          if(count === 3){
            that.element.remove()
            return
          }
        }
        idx %= 15
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.hit = function(_beAtker){
    let diffTime = 0
    let lastDate = Date.now()
    let idx = 0
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        img.src = require('@/assets/frameimg/skills/1003/hit_' + idx + '.png')
        img.onload = (() => {
          that.hitcav.clearRect(0, 0, that.hitelement.width, that.hitelement.height)
          that.hitcav.drawImage(img, _beAtker.xStand, _beAtker.yStand - 2.5, 100, 101)
        })
        idx++
        if(idx === 14){
          that.hitelement.remove()
          return
        }
        idx %= 14
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.doSecond = function(_atker, _beAtker){
    this.skillStatus = 1
    let xOffset = -150
    let count = 0
    if(_beAtker.type == 1){
      // 1表示我被攻击
      this.cav2.scale(-1, 1);
      this.cav2.translate(-this.element2.width / this.dpr, 0);
      xOffset = 30
    }
    let diffTime = 0
    let lastDate = Date.now()
    let idx = 0
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        img.src = require('@/assets/frameimg/skills/1003/cast_' + idx + '.png')
        img.onload = (() => {
          that.cav2.clearRect(0, 0, that.element2.width, that.element2.height)
          that.cav2.drawImage(img, _beAtker.xStand + xOffset, _beAtker.yStand + 7, 143, 80)
        })
        idx++
        xOffset = xOffset + 8
        if(idx === 15){
          count++
          if(count === 1){
            that.hitSecond(_beAtker)
          }
          if(count === 3){
            that.skillStatus = 99
            that.element2.remove()
            return
          }
        }
        idx %= 15
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.hitSecond = function(_beAtker){
    let diffTime = 0
    let lastDate = Date.now()
    let idx = 0
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        img.src = require('@/assets/frameimg/skills/1003/hit_' + idx + '.png')
        img.onload = (() => {
          that.hitcav2.clearRect(0, 0, that.hitelement2.width, that.hitelement2.height)
          that.hitcav2.drawImage(img, _beAtker.xStand, _beAtker.yStand - 2.5, 100, 101)
        })
        idx++
        if(idx === 14){
          that.hitelement2.remove()
          return
        }
        idx %= 14
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }
}
