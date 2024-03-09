import { createPVPCavElement, createPVPCav } from '@/animation/cavUtil'

export function Sk1002(_container, _dpr){
  this.skillStatus = 0 // 0 init  1 doing  99 finish
  this.dpr = _dpr
  this.element = createPVPCavElement(_container, _dpr, true, 0)
  this.cav = createPVPCav(this.element, _dpr)
  this.hitelement = createPVPCavElement(_container, _dpr, true, 0)
  this.hitcav = createPVPCav(this.hitelement, _dpr)

  this.do = function(_atker, _beAtker){
    this.skillStatus = 1
    let diffTime = 0
    let lastDate = Date.now()
    let idx = 0
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        img.src = require('@/assets/frameimg/skills/1002/cast_' + idx + '.png')
        img.onload = (() => {
          that.cav.clearRect(0, 0, that.element.width, that.element.height)
          that.cav.drawImage(img, _beAtker.xStand, _beAtker.yStand - 30, 100, 166)
        })
        idx++
        if(idx === 33){
          that.hit(_beAtker)
        }
        if(idx === 35){
          that.element.remove()
          return
        }
        idx %= 35
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.hit = function(_beAtker){
    let diffTime = 20
    let lastDate = Date.now()
    let idx = 0
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        img.src = require('@/assets/frameimg/skills/1002/hit_' + idx + '.png')
        img.onload = (() => {
          that.hitcav.clearRect(0, 0, that.hitelement.width, that.hitelement.height)
          that.hitcav.drawImage(img, _beAtker.xStand - 20, _beAtker.yStand - 66, 140, 180)
        })
        idx++
        if(idx === 41){
          that.skillStatus = 99
          that.hitelement.remove()
          return
        }
        idx %= 41
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

}