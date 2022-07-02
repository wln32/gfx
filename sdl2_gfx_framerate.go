package gfx

/*

	#include "SDL2_framerate.h"


	Uint32 _getTicks()
	{
	Uint32 ticks = SDL_GetTicks();


	if (ticks == 0) {
	return 1;
	} else {
	return ticks;
	}
	}


	void SDL_initFramerate(FPSmanager * manager)
	{

	manager->framecount = 0;
	manager->rate = FPS_DEFAULT;
	manager->rateticks = (1000.0f / (float) FPS_DEFAULT);
	manager->baseticks = _getTicks();
	manager->lastticks = manager->baseticks;

	}


	int SDL_setFramerate(FPSmanager * manager, Uint32 rate)
	{
	if ((rate >= FPS_LOWER_LIMIT) && (rate <= FPS_UPPER_LIMIT)) {
	manager->framecount = 0;
	manager->rate = rate;
	manager->rateticks = (1000.0f / (float) rate);
	return (0);
	} else {
	return (-1);
	}
	}


	int SDL_getFramerate(FPSmanager * manager)
	{
	if (manager == NULL) {
	return (-1);
	} else {
	return ((int)manager->rate);
	}
	}


	int SDL_getFramecount(FPSmanager * manager)
	{
	if (manager == NULL) {
	return (-1);
	} else {
	return ((int)manager->framecount);
	}
	}


Uint32 SDL_framerateDelay(FPSmanager * manager)
{
	Uint32 current_ticks;
	Uint32 target_ticks;
	Uint32 the_delay;
	Uint32 time_passed = 0;


	if (manager == NULL) {
	return 0;
	}


	if (manager->baseticks == 0) {
	SDL_initFramerate(manager);
	}


	manager->framecount++;


	current_ticks = _getTicks();
	time_passed = current_ticks - manager->lastticks;
	manager->lastticks = current_ticks;
	target_ticks = manager->baseticks + (Uint32) ((float) manager->framecount * manager->rateticks);

	if (current_ticks <= target_ticks) {
	the_delay = target_ticks - current_ticks;
	SDL_Delay(the_delay);
	} else {
	manager->framecount = 0;
	manager->baseticks = _getTicks();
	}

	return time_passed;
}
*/
import "C"
