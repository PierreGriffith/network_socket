
	var socket;
	var curr_player;
	var x,y;
	var players = [];
	var zoom = 1;


function setup() 
{
	createCanvas(600, 600);


	
	socket = io();


	curr_player = new player(random(width), random(height), 0, 10) 


	socket.on('position',
    function(data) {
    	players = JSON.parse(data)      	   
    });


}

function draw()
{	
	background(0);
	translate(width / 2, height / 2);
  	var newzoom = 64 / curr_player.r;
  	zoom = lerp(zoom, newzoom, 0.1);
  	scale(zoom);
  	translate(-curr_player.pos.x, -curr_player.pos.y);

	/*
	if (players.length != 0) 
	{
		for (var i =  0; i != players.length -1; i++ ) 
		{
			if (players.id !== socket.id)
			{				
				fill(0, 0, 255);         	
				ellipse(players[i].x, players[i].y ,80,80);			
    		}
    	}		
	}
	*/
 if(mouseIsPressed)
    camera.zoom = .5;
  else
    camera.zoom = 1;

	curr_player.show();
	curr_player.update();    	
	curr_player.constrain();
	console.log(players)

	socket.emit('position',JSON.stringify({"x" :curr_player.pos.x, "y": curr_player.pos.y, "id" : socket.id }));			
	
}