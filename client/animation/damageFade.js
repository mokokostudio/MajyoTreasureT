export function Damage(_x, _y, _damage,_cavCtx, _element){
  this.x = _x
  this.y = _y
  this.damage = _damage === 0 ? 'MISS' : _damage
  this.canvas = _cavCtx
  this.element = _element

  this.xSpd = 0
  this.ySpd = 0
  this.opacity = 1
  this.size = 1
  this.sizeMuti = 1
  this.direction = 0

  this.moveDirection = function(_dir){
    this.direction = _dir
  }

  // 文字颜色RGB 描边（92,33,2）  白色（255,244,208），黄色（255,211,70）

  this.fadeNormalDmgToBoss = () =>{
    this.canvas.font = "600 italic 16px myFirstFont";
    this.canvas.textBaseline = 'alphabetic';
    this.canvas.fillStyle = "rgba(255,244,208,"+this.opacity+")";
    this.canvas.strokeStyle = "rgba(92,33,2,"+this.opacity+")"; // 描边颜色
    this.canvas.lineWidth = 2; // 描边宽度
    this.canvas.clearRect(0, 0, this.element.width, this.element.height);
    this.canvas.strokeText(this.damage, this.x + this.xSpd, this.y + this.ySpd)
    this.canvas.fillText(this.damage, this.x + this.xSpd, this.y + this.ySpd)
    this.ySpd = this.ySpd - 1
    this.xSpd = this.xSpd + 1
    this.opacity = this.opacity - 0.02
    if(this.opacity <= 0){
      this.element.remove()
      return
    }
    requestAnimationFrame(this.fadeNormalDmgToBoss)
  }

  this.fadeCriDmgToBoss = () =>{
    this.canvas.font = "600 italic " + 20 * this.size + "px myFirstFont";
    this.canvas.textBaseline = 'alphabetic';
    this.canvas.fillStyle = "rgba(255,211,70," + this.opacity + ")";
    this.canvas.strokeStyle = "rgba(92,33,2," + this.opacity + ")"; // 描边颜色
    this.canvas.lineWidth = 2; // 描边宽度
    this.canvas.clearRect(0, 0, this.element.width, this.element.height);
    this.canvas.strokeText(this.damage, this.x + this.xSpd, this.y + this.ySpd)
    this.canvas.fillText(this.damage, this.x + this.xSpd, this.y + this.ySpd)
    this.size = this.size + 0.25 * this.sizeMuti
    if(this.size >= 1.5){
      this.sizeMuti = -1
    }
    if(this.size <= 1){
      this.sizeMuti = 0
    }
    this.ySpd = this.ySpd - 1
    this.opacity = this.opacity - 0.02
    if(this.opacity <= 0){
      this.element.remove()
      return
    }
    requestAnimationFrame(this.fadeCriDmgToBoss)
  }

  this.fadeNormalDmgToPVP = () =>{
    this.canvas.font = "600 italic 14px myFirstFont";
    this.canvas.textBaseline = 'alphabetic';
    this.canvas.fillStyle = "rgba(255,244,208," + this.opacity + ")";
    this.canvas.strokeStyle = "rgba(92,33,2," + this.opacity + ")";
    this.canvas.lineWidth = 2;
    this.canvas.clearRect(0, 0, this.element.width, this.element.height);
    this.canvas.strokeText(this.damage, this.x + this.xSpd, this.y + this.ySpd)
    this.canvas.fillText(this.damage, this.x + this.xSpd, this.y + this.ySpd)
    this.ySpd = this.ySpd - 0.5
    this.opacity = this.opacity - 0.025
    if(this.direction === 1){
      // 往右
      this.xSpd = this.xSpd + 1
    }else{
      // 往左
      this.xSpd = this.xSpd - 1
    }
    if(this.opacity <= 0){
      this.element.remove()
      return
    }
    requestAnimationFrame(this.fadeNormalDmgToPVP)
  }

  this.fadeCriDmgToPVP = () =>{
    this.canvas.font = "600 italic " + 18 * this.size + "px myFirstFont";
    this.canvas.textBaseline = 'alphabetic';
    this.canvas.fillStyle = "rgba(255,211,70," + this.opacity + ")";
    this.canvas.strokeStyle = "rgba(92,33,2," + this.opacity + ")";
    this.canvas.lineWidth = 2;
    this.canvas.clearRect(0, 0, this.element.width, this.element.height);
    this.canvas.strokeText(this.damage, this.x + this.xSpd, this.y + this.ySpd)
    this.canvas.fillText(this.damage, this.x + this.xSpd, this.y + this.ySpd)
    this.size = this.size + 0.5 * this.sizeMuti
    if(this.size >= 1.5){
      this.sizeMuti = -1
    }
    if(this.size <= 1){
      this.sizeMuti = 0
    }
    this.ySpd = this.ySpd - 0.5
    this.opacity = this.opacity - 0.025
    if(this.direction === 1){
      // 往右
      this.xSpd = this.xSpd + 1
    }else{
      // 往左
      this.xSpd = this.xSpd - 1
    }
    if(this.opacity <= 0){
      this.element.remove()
      return
    }
    requestAnimationFrame(this.fadeCriDmgToPVP)
  }
}
