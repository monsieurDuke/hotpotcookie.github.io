(function(window, document, undefined){

	window.onload = init;
	function init(){
		// Button Path Marker
		// ------------------
		const cyan_old = "#ab87e6";
		const pink = "#FF8AB8";
		const white = "#FFF";
		const yellow = "#FFEC9D";

		const sidenav_button = [
			document.getElementById('side-home'),		
			document.getElementById('side-tutorial'),
			document.getElementById('side-walktrough'),
			document.getElementById('side-publication'),
			document.getElementById('side-profile')			
			// button bisa scalable
		]
		const topnav_button = [
			document.getElementById('top-home'),		
			document.getElementById('top-tutorial'),
			document.getElementById('top-walktrough'),
			document.getElementById('top-publication'),
			document.getElementById('top-profile')			
			// button bisa scalable
		]

		var idx_navbutton = 0;
		const current_URL = window.location.href.toString().split(window.location.host)[1];
		console.log(current_URL);

		switch(true) {
			case current_URL.indexOf("tutorial") !== -1:
				idx_navbutton = 1;
				break;			
			case current_URL.indexOf("walktrough") !== -1:
				idx_navbutton = 2;
				break;						
			case current_URL.indexOf("publication") !== -1:
				idx_navbutton = 3;
				break;
			case current_URL.indexOf("profile") !== -1:
				idx_navbutton = 4;
				break;
			case current_URL.indexOf("privacy-policy") !== -1:
				idx_navbutton = -1;
				break;				
			default:
				idx_navbutton = 0;
				break;		
		}
		for (var i=0; i<sidenav_button.length; i++) {
			if (i != idx_navbutton) {
				sidenav_button[i].style.color = pink;
				topnav_button[i].style.color = white;
				console.log(i+" no");
			} else {
				sidenav_button[i].style.color = cyan_old;
				topnav_button[i].style.color = yellow;				
				console.log(i+" yes");				
			}
		}		
	}
})(window, document, undefined);

function toggleDesc(x) {
	let arr = x.id.split("-");
	key_class = "text-"+arr[1];	
	if (x.className.includes("up")) {
		x.className = x.className.replace("up","down");
		document.getElementById(key_class).style.display = 'flex';
	} else {
		console.log(x.className);
		x.className = x.className.replace("down","up");
		document.getElementById(key_class).style.display = 'none';		
	}
}

