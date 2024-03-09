export function fzShadow(){

  this.show = function(_skillCode, _character){
    if(_skillCode === '1004'){
      this.doShadowFront('1001', 10, 8, _character)
    }else if(_skillCode === '1000' || _skillCode === '1001'){
      this.doShadow('1002', 12, 10, _character)
    }else if(_skillCode === '1002' || _skillCode === '1003' || _skillCode === '1005' || _skillCode === '1008'){
      this.doShadow('1001', 12, 10, _character)
    }else if(_skillCode === '1009'){
      this.doShadowFront('1003', 10, 8, _character)
    }else{
      this.doShadow('1000', 12, 10, _character)
    }
  }

  this.doShadow = function(_fzcode, _frame, _scaleGap, _character){
    let idx = 0
    let scalediff = 0
    let xdiff = _character.type == 1 ? 10 : -10
    let xpos= _character.xStand + _character.mskAtr.width / 2
    let ypos = _character.yStand + _character.mskAtr.height
    let diffTime = 0
    let lastDate = Date.now()
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let img = new Image()
        img.src = require('@/assets/frameimg/fz/' + _fzcode + '.png')
        img.onload = (() => {
          _character.shadowCav.clearRect(0, 0, _character.shadowElement.width, _character.shadowElement.height)
          _character.shadowCav.drawImage(img, xpos + xdiff, ypos, scalediff, scalediff * 0.4)

        })
        scalediff += _scaleGap
        xpos -= _scaleGap / 2
        ypos -= _scaleGap * 0.2

        idx++
        if(idx === _frame){
          return
        }
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.doShadowFront = function(_fzcode, _frame, _scaleGap, _character){
    let idx = 0
    let scalediff = 0
    let xpos = _character.type == 1 ? _character.xStand + _character.mskAtr.width + 20 : _character.xStand - 20
    let ypos = _character.yStand + _character.mskAtr.height / 2
    let diffTime = 0
    let lastDate = Date.now()
    let doDraw = (() =>{
      let curDate = Date.now()
      if(curDate - lastDate >= diffTime){
        lastDate = curDate
        let img = new Image()
        img.src = require('@/assets/frameimg/fz/' + _fzcode + '.png')
        img.onload = (() => {
          _character.shadowCav.clearRect(0, 0, _character.shadowElement.width, _character.shadowElement.height)
          _character.shadowCav.drawImage(img, xpos, ypos, scalediff * 0.4, scalediff)

        })
        scalediff += _scaleGap
        xpos -= _scaleGap * 0.2
        ypos -= _scaleGap / 2

        idx++
        if(idx === _frame){
          return
        }
      }
      requestAnimationFrame(doDraw)
    })
    requestAnimationFrame(doDraw)
  }

  this.recoverShadow = function(_character){
    let shadow = new Image()
    shadow.src = require('@/assets/frameimg/shadow.png')
    shadow.onload = (() => {
      _character.shadowCav.clearRect(0, 0, _character.shadowElement.width, _character.shadowElement.height)
      _character.shadowCav.drawImage(shadow, _character.shadowXpos, _character.shadowYpos, 80, 27)
    })
  }
}


