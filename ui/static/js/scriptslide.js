var btn_prev=document.querySelector('#gallery .buttons .prev');
var btn_next=document.querySelector('#gallery .buttons .next');
var images=document.querySelectorAll('#gallery .photos img');
console.log(images);
var i=0;
btn_prev.onclick = function() {
	/*images[i].style.display='none';*/
	images[i].className='';
	i--;
	if(i<0){
		i=images.length-1;
	}
	/*images[i].style.display='block';*/
	images[i].className='Showed';
}
btn_next.onclick = function() {
	/*images[i].style.display='none';*/
	images[i].className='';
	i++;
	if(i>=images.length){
		i=0;
	}
	images[i].className='Showed';
	/*images[i].style.display='block';*/
}