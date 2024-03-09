import { fzShadow } from '@/animation/shadowUtil'
import { Skill } from '@/animation/skill'

export function MainCharacter(_type, _dpr, _cavsContainer){
  this.type = _type // 1 自己  2 对手
  this.dpr = _dpr
  this.container = _cavsContainer
  this.hpLeft = 0
  this.skillAni = new Skill(_cavsContainer, _dpr)
  this.fzShadow = new fzShadow()

  this.mainElement = document.createElement('canvas');
  this.mainElement.style.width = "380px";
  this.mainElement.style.height = "230px";
  this.mainElement.className  = this.type === 1 ? 'battlecav' : 'battlecav'
  this.mainElement.width = 380 * this.dpr;
  this.mainElement.height = 230 * this.dpr;
  this.mainElement.style.zIndex = this.type === 1 ? 50 : 100
  this.container.appendChild(this.mainElement);
  this.mainCav = this.mainElement.getContext('2d');
  this.mainCav.scale(this.dpr, this.dpr)

  this.shadowElement = document.createElement('canvas');
  this.shadowElement.style.width = "380px";
  this.shadowElement.style.height = "230px";
  this.shadowElement.className  = 'battlecav'
  this.shadowElement.width = 380 * this.dpr;
  this.shadowElement.height = 230 * this.dpr;
  this.shadowElement.style.zIndex = this.type === 1 ? 10 : 11
  this.container.appendChild(this.shadowElement);
  this.shadowCav = this.shadowElement.getContext('2d');
  this.shadowCav.scale(this.dpr, this.dpr)

  // 0 init  99 finish  98 die -98 dieing
  // 1 move   2 stand  3 atk  4 beAtk  -4 beatking  5 skill
  this.status = 0
  this.beSkill = 0 // 0   1 beskill
  this.isWin = 0 // 0 初始化  1 是  2 否

  this.width = 70
  this.height = 84
  this.resource = _type === 1 ? '' : '2'

  this.moveAtr = {
    width: 74,
    height: 93,
    direction : this.type === 1 ? 1 : -1,
    xStart : this.type === 1 ? 0 : this.mainElement.width / this.dpr - 74,
    yStart : this.mainElement.height/ this.dpr / 2 - 93 / 2 - 10
  }

  this.atkAtr = {
    width: 95,
    height: 97,
    shadowDir : this.type === 1 ? 1 : -1
  }

  this.beatkAtr = {
    width: 78,
    height: 90
  }

  this.dieAtr = {
    width: 78,
    height: 91
  }

  this.mskAtr = {
    width: 68,
    height: 83
  }

  this.xStand = this.moveAtr.xStart
  this.yStand = this.moveAtr.yStart
  this.shadowXpos = 0
  this.shadowYpos = 0

  // 监听是否受击
  let beAtkTimer = setInterval(() => {
    console.log('listen be atk', this.type, this.status)
    if(this.status == 4){
      this.beAttack()
    }else if(this.status == 99){
      clearInterval(beAtkTimer)
    }
  },100)

  // 监听是否死亡
  let dieTimer = setInterval(() => {
    console.log('listen die', this.type, this.status)
    if(this.status == 98){
      this.die()
    }else if(this.status == 99){
      clearInterval(dieTimer)
    }
  },100)

  this.stand = function(fromMove){
    console.log('stand==', this.type)
    if(fromMove){
      this.yStand -= 5
    }
    let wait = 0
    let idx = 0
    let diffTime = 0
    let lastDate = Date.now()
    let doDraw = (() =>{
      //console.log('type ',this.type," status ", this.status)
      if(wait == 5 && this.isWin == 1){
        this.status = 99
        return
      }
      //console.log("stand ",this.xStand, this.yStand)
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        img.src = require('@/assets/frameimg/main/idle'+that.resource+'_' + idx + '.png')
        img.onload = (() => {
          that.mainCav.clearRect(0, 0, that.mainElement.width, that.mainElement.height)
          that.mainCav.drawImage(img, that.xStand, that.yStand, that.width, that.height)
        })
        idx++
        if(idx === 1){
          diffTime = 120
        }
        if(idx === 16){
          wait++
        }
        if(this.status !== 2){
          return
        }
        idx %= 16
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.move = function(_fps){
    console.log('move==', this.type)
    this.status = 1
    let idx = 0
    let movediff = 0
    let diffTime = _fps; // 动画最小时间间隔,单位 ms
    let lastDate = Date.now()
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        console.log(idx)
        img.src = require('@/assets/frameimg/main/move'+that.resource+'_' + idx + '.png')
        img.onload = (() => {
          that.mainCav.clearRect(0, 0, that.mainElement.width, that.mainElement.height)
          that.mainCav.drawImage(img, that.xStand, that.yStand, that.moveAtr.width, that.moveAtr.height)

          let shadow = new Image()
          shadow.src = require('@/assets/frameimg/shadow.png')
          shadow.onload = (() => {
            let xpos = that.type === 1 ? 6 : -15
            that.shadowXpos = that.xStand + xpos
            that.shadowYpos = that.yStand + 64
            that.shadowCav.clearRect(0, 0, that.shadowElement.width, that.shadowElement.height)
            that.shadowCav.drawImage(shadow, that.shadowXpos, that.shadowYpos, 80, 27)

            if(idx === 16){
              that.status = 2
              that.stand(true)
            }
          })
        })

        that.xStand = that.moveAtr.xStart + movediff * that.moveAtr.direction
        movediff += 4
        idx++
        if(idx === 16){
          return
        }
        idx %= 16
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.attack = function(_fps, _character){
    console.log('atk==', this.type)
    this.status = 3
    let idx = 0
    let xshadow = 0
    let xdiff = this.type === 1 ? -7: -18
    let diffTime = _fps;
    let lastDate = Date.now()
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        //console.log('atk=',idx, 'ystand', that.yStand)
        img.src = require('@/assets/frameimg/main/atk'+that.resource+'_' + idx + '.png')
        img.onload = (() => {
          that.mainCav.clearRect(0, 0, that.mainElement.width, that.mainElement.height)
          that.mainCav.drawImage(img, that.xStand + xdiff, that.yStand, that.atkAtr.width, that.atkAtr.height)

          if(idx >= 4 && idx <= 7){
            xshadow += 1
          }else if(idx >= 12 && idx <= 15){
            xshadow -= 1
          }
          let shadow = new Image()
          shadow.src = require('@/assets/frameimg/shadow.png')
          shadow.onload = (() => {
            that.shadowCav.clearRect(0, 0, that.shadowElement.width, that.shadowElement.height)
            that.shadowCav.drawImage(shadow, that.shadowXpos + xshadow * that.atkAtr.shadowDir, that.shadowYpos, 80, 27)
          })
        })

        idx++
        if(idx == 5){
          if(_character.hpLeft == 0){
            _character.status = 98
            that.isWin = 1
          }else{
            _character.status = 4
          }
        }
        if(idx == 16){
          console.log('winner',this.type,this)
          that.status = 2
          that.stand(false)
          return
        }
        idx %= 16
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.beAttack = function(){
    console.log('be atk==', this.type)
    this.status = -4
    let idx = 0
    let xdiff = this.type === 1 ? -2: -6
    let diffTime = 0;
    let lastDate = Date.now()
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        img.src = require('@/assets/frameimg/main/beatk'+that.resource+'_' + idx + '.png')
        img.onload = (() => {
          that.mainCav.clearRect(0, 0, that.mainElement.width, that.mainElement.height)
          that.mainCav.drawImage(img, that.xStand + xdiff, that.yStand, that.beatkAtr.width, that.beatkAtr.height)
        })
        // if(idx == 0){
        //   console.log(this)
        //   return
        // }
        idx++
        if(idx === 1){
          if(that.beSkill == 0){
            that.beAttackEft()
          }
        }
        if(idx == 16){
          that.status = 2
          that.beSkill = 0
          that.stand(false)
          return
        }
        idx %= 16
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.die = function(){
    console.log('die==', this.type)
    this.status = -98
    this.isWin = 2
    let idx = 0
    let xdiff = this.type === 1 ? 1: -9
    let diffTime = 20;
    let lastDate = Date.now()
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        img.src = require('@/assets/frameimg/main/die'+that.resource+'_' + idx + '.png')
        img.onload = (() => {
          that.mainCav.clearRect(0, 0, that.mainElement.width, that.mainElement.height)
          that.mainCav.drawImage(img, that.xStand + xdiff, that.yStand, that.dieAtr.width, that.dieAtr.height)
        })
        idx++
        if(idx === 1){
          if(that.beSkill == 0){
            that.beAttackEft()
          }
        }
        if(idx == 28){
          that.status = 99
          that.beSkill = 0
          return
        }
        idx %= 28
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.beAttackEft = function(){
    console.log('be atk eft==', this.type)
    let _element = document.createElement('canvas');
    _element.style.width = "380px";
    _element.style.height = "230px";
    _element.className  = 'battlecav'
    _element.width = 380 * this.dpr;
    _element.height = 230 * this.dpr;
    let t = new Date().getTime().toString()
    t = parseInt(t.substring(t.length - 6, t.length))
    _element.style.zIndex = t
    this.container.appendChild(_element);
    let _cav = _element.getContext('2d');
    _cav.scale(this.dpr, this.dpr)
    let idx = 0
    let xdiff = this.type == 1 ? -25 : -50
    let doDraw = (() =>{
      let that = this
      let img = new Image()
      img.src = require('@/assets/frameimg/be_atk/be-atk_' + idx + '.png')
      img.onload = (() => {
        _cav.clearRect(0, 0, _element.width, _element.height)
        _cav.drawImage(img, that.xStand + that.width / 2 + xdiff, that.yStand + that.height / 2 - 30, 80, 80)
      })
      idx++
      if(idx == 13){
        _element.remove()
        return
      }
      idx %= 13
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.Skill = function(_skillCode, _beAtker){
    console.log('do sk pre==', this.type)
    this.status = 5
    let idx = 0
    let diffTime = 0;
    let lastDate = Date.now()
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        //console.log('atk=',idx, 'ystand', that.yStand)
        img.src = require('@/assets/frameimg/main/msk'+that.resource+'_' + idx + '.png')
        img.onload = (() => {
          that.mainCav.clearRect(0, 0, that.mainElement.width, that.mainElement.height)
          that.mainCav.drawImage(img, that.xStand, that.yStand, that.mskAtr.width, that.mskAtr.height)
        })
        idx++
        if(idx === 4){
          that.fzShadow.show(_skillCode, that)
        }
        if(idx == 11){
          that.skillCast(_skillCode, _beAtker)
          that.skillAni.cast(_skillCode, that, _beAtker)
          return
        }
        idx %= 11
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.skillCast = function(_skillCode, _beAtker){
    console.log('do sk cast==', this.type)
    let idx = 12
    let diffTime = 51;
    let count = 0
    let lastDate = Date.now()
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        //console.log('atk=',idx, 'ystand', that.yStand)
        img.src = require('@/assets/frameimg/main/msk'+that.resource+'_' + idx + '.png')
        img.onload = (() => {
          that.mainCav.clearRect(0, 0, that.mainElement.width, that.mainElement.height)
          that.mainCav.drawImage(img, that.xStand, that.yStand, that.mskAtr.width, that.mskAtr.height)
        })
        idx++
        idx %= 21
        if(idx === 0){
          count ++
          idx = 12
        }
        if(count == 0 && idx == 15){
          if(_beAtker.hpLeft == 0){
            _beAtker.beSkill = 1
            _beAtker.status = 98
            that.isWin = 1
          }
        }
        if(that.skillAni.status == 99){
          that.finishSkill()
          return
        }
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }


  this.finishSkill = function(){
    console.log('finish sk==', this.type)
    let idx = 21
    let diffTime = 0;
    let lastDate = Date.now()
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let that = this
        let img = new Image()
        //console.log('atk=',idx, 'ystand', that.yStand)
        img.src = require('@/assets/frameimg/main/msk'+that.resource+'_' + idx + '.png')
        img.onload = (() => {
          that.mainCav.clearRect(0, 0, that.mainElement.width, that.mainElement.height)
          that.mainCav.drawImage(img, that.xStand, that.yStand, that.mskAtr.width, that.mskAtr.height)
        })
        idx++
        if(idx === 26){
          that.fzShadow.recoverShadow(that)
          that.status = 2
          that.stand(false)
          return
        }
        idx %= 26
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

}
