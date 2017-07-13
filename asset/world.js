


function player(x, y, z, r) {

	this.pos = createVector(x, y, z)
	this.r = r
	this.vel = createVector(0,0)

   this.update = function() 
   {
 
    var newvel = createVector(mouseX - width / 2, mouseY - height / 2); // vecteur longueur
  //  newvel.div(50);
    newvel.limit(3); // longueur max
    this.vel.lerp(newvel, 0.2);
    this.pos.add(this.vel);
	
	}

   this.constrain = function() {
    this.pos.x = constrain(this.pos.x, -width / 4, width / 4);
    this.pos.y = constrain(this.pos.y, -height / 4, height / 4);
  }

	this.show = function () {
    fill(255);
    ellipse(this.pos.x, this.pos.y, this.r * 2, this.r * 2);	
	}

}