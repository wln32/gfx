package gfx

/*
#ifdef WIN32
#include <windows.h>
#endif

#include <stdlib.h>
#include <string.h>

#include "SDL2_rotozoom.h"




typedef struct tColorRGBA {
	Uint8 r;
	Uint8 g;
	Uint8 b;
	Uint8 a;
} tColorRGBA;


typedef struct tColorY {
	Uint8 y;
} tColorY;


#define MAX(a,b)    (((a) > (b)) ? (a) : (b))


#define GUARD_ROWS (2)


#define VALUE_LIMIT	0.001


Uint32 _colorkey(SDL_Surface *src)
{
	Uint32 key = 0;
	SDL_GetColorKey(src, &key);
	return key;
}



int _shrinkSurfaceRGBA(SDL_Surface * src, SDL_Surface * dst, int factorx, int factory)
{
	int x, y, dx, dy, dgap, ra, ga, ba, aa;
	int n_average;
	tColorRGBA *sp, *osp, *oosp;
	tColorRGBA *dp;




	n_average = factorx*factory;


	sp = (tColorRGBA *) src->pixels;

	dp = (tColorRGBA *) dst->pixels;
	dgap = dst->pitch - dst->w * 4;

	for (y = 0; y < dst->h; y++) {

		osp=sp;
		for (x = 0; x < dst->w; x++) {


			oosp=sp;
			ra=ga=ba=aa=0;
			for (dy=0; dy < factory; dy++) {
				for (dx=0; dx < factorx; dx++) {
					ra += sp->r;
					ga += sp->g;
					ba += sp->b;
					aa += sp->a;

					sp++;
				}

				sp = (tColorRGBA *)((Uint8*)sp + (src->pitch - 4*factorx)); // next y
			}



			sp = (tColorRGBA *)((Uint8*)oosp + 4*factorx);


			dp->r = ra/n_average;
			dp->g = ga/n_average;
			dp->b = ba/n_average;
			dp->a = aa/n_average;


			dp++;
		}



		sp = (tColorRGBA *)((Uint8*)osp + src->pitch*factory);


		dp = (tColorRGBA *) ((Uint8 *) dp + dgap);
	}


	return (0);
}


int _shrinkSurfaceY(SDL_Surface * src, SDL_Surface * dst, int factorx, int factory)
{
	int x, y, dx, dy, dgap, a;
	int n_average;
	Uint8 *sp, *osp, *oosp;
	Uint8 *dp;




	n_average = factorx*factory;


	sp = (Uint8 *) src->pixels;

	dp = (Uint8 *) dst->pixels;
	dgap = dst->pitch - dst->w;

	for (y = 0; y < dst->h; y++) {

		osp=sp;
		for (x = 0; x < dst->w; x++) {


			oosp=sp;
			a=0;
			for (dy=0; dy < factory; dy++) {
				for (dx=0; dx < factorx; dx++) {
					a += (*sp);

					sp++;
				}


				sp = (Uint8 *)((Uint8*)sp + (src->pitch - factorx));
			}



			sp = (Uint8 *)((Uint8*)oosp + factorx);


			*dp = a/n_average;


			dp++;
		}



		sp = (Uint8 *)((Uint8*)osp + src->pitch*factory);


		dp = (Uint8 *)((Uint8 *)dp + dgap);
	}


	return (0);
}


int _zoomSurfaceRGBA(SDL_Surface * src, SDL_Surface * dst, int flipx, int flipy, int smooth)
{
	int x, y, sx, sy, ssx, ssy, *sax, *say, *csax, *csay, *salast, csx, csy, ex, ey, cx, cy, sstep, sstepx, sstepy;
	tColorRGBA *c00, *c01, *c10, *c11;
	tColorRGBA *sp, *csp, *dp;
	int spixelgap, spixelw, spixelh, dgap, t1, t2;


	if ((sax = (int *) malloc((dst->w + 1) * sizeof(Uint32))) == NULL) {
		return (-1);
	}
	if ((say = (int *) malloc((dst->h + 1) * sizeof(Uint32))) == NULL) {
		free(sax);
		return (-1);
	}


	spixelw = (src->w - 1);
	spixelh = (src->h - 1);
	if (smooth) {
		sx = (int) (65536.0 * (float) spixelw / (float) (dst->w - 1));
		sy = (int) (65536.0 * (float) spixelh / (float) (dst->h - 1));
	} else {
		sx = (int) (65536.0 * (float) (src->w) / (float) (dst->w));
		sy = (int) (65536.0 * (float) (src->h) / (float) (dst->h));
	}


	ssx = (src->w << 16) - 1;
	ssy = (src->h << 16) - 1;


	csx = 0;
	csax = sax;
	for (x = 0; x <= dst->w; x++) {
		*csax = csx;
		csax++;
		csx += sx;


		if (csx > ssx) {
			csx = ssx;
		}
	}


	csy = 0;
	csay = say;
	for (y = 0; y <= dst->h; y++) {
		*csay = csy;
		csay++;
		csy += sy;


		if (csy > ssy) {
			csy = ssy;
		}
	}

	sp = (tColorRGBA *) src->pixels;
	dp = (tColorRGBA *) dst->pixels;
	dgap = dst->pitch - dst->w * 4;
	spixelgap = src->pitch/4;

	if (flipx) sp += spixelw;
	if (flipy) sp += (spixelgap * spixelh);


	if (smooth) {


		csay = say;
		for (y = 0; y < dst->h; y++) {
			csp = sp;
			csax = sax;
			for (x = 0; x < dst->w; x++) {

				ex = (*csax & 0xffff);
				ey = (*csay & 0xffff);
				cx = (*csax >> 16);
				cy = (*csay >> 16);
				sstepx = cx < spixelw;
				sstepy = cy < spixelh;
				c00 = sp;
				c01 = sp;
				c10 = sp;
				if (sstepy) {
					if (flipy) {
						c10 -= spixelgap;
					} else {
						c10 += spixelgap;
					}
				}
				c11 = c10;
				if (sstepx) {
					if (flipx) {
						c01--;
						c11--;
					} else {
						c01++;
						c11++;
					}
				}


				t1 = ((((c01->r - c00->r) * ex) >> 16) + c00->r) & 0xff;
				t2 = ((((c11->r - c10->r) * ex) >> 16) + c10->r) & 0xff;
				dp->r = (((t2 - t1) * ey) >> 16) + t1;
				t1 = ((((c01->g - c00->g) * ex) >> 16) + c00->g) & 0xff;
				t2 = ((((c11->g - c10->g) * ex) >> 16) + c10->g) & 0xff;
				dp->g = (((t2 - t1) * ey) >> 16) + t1;
				t1 = ((((c01->b - c00->b) * ex) >> 16) + c00->b) & 0xff;
				t2 = ((((c11->b - c10->b) * ex) >> 16) + c10->b) & 0xff;
				dp->b = (((t2 - t1) * ey) >> 16) + t1;
				t1 = ((((c01->a - c00->a) * ex) >> 16) + c00->a) & 0xff;
				t2 = ((((c11->a - c10->a) * ex) >> 16) + c10->a) & 0xff;
				dp->a = (((t2 - t1) * ey) >> 16) + t1;

				salast = csax;
				csax++;
				sstep = (*csax >> 16) - (*salast >> 16);
				if (flipx) {
					sp -= sstep;
				} else {
					sp += sstep;
				}


				dp++;
			}

			salast = csay;
			csay++;
			sstep = (*csay >> 16) - (*salast >> 16);
			sstep *= spixelgap;
			if (flipy) {
				sp = csp - sstep;
			} else {
				sp = csp + sstep;
			}


			dp = (tColorRGBA *) ((Uint8 *) dp + dgap);
		}
	} else {

		csay = say;
		for (y = 0; y < dst->h; y++) {
			csp = sp;
			csax = sax;
			for (x = 0; x < dst->w; x++) {

				*dp = *sp;


				salast = csax;
				csax++;
				sstep = (*csax >> 16) - (*salast >> 16);
				if (flipx) sstep = -sstep;
				sp += sstep;


				dp++;
			}

			salast = csay;
			csay++;
			sstep = (*csay >> 16) - (*salast >> 16);
			sstep *= spixelgap;
			if (flipy) sstep = -sstep;
			sp = csp + sstep;


			dp = (tColorRGBA *) ((Uint8 *) dp + dgap);
		}
	}


	free(sax);
	free(say);

	return (0);
}


int _zoomSurfaceY(SDL_Surface * src, SDL_Surface * dst, int flipx, int flipy)
{
	int x, y;
	Uint32 *sax, *say, *csax, *csay;
	int csx, csy;
	Uint8 *sp, *dp, *csp;
	int dgap;


	if ((sax = (Uint32 *) malloc((dst->w + 1) * sizeof(Uint32))) == NULL) {
		return (-1);
	}
	if ((say = (Uint32 *) malloc((dst->h + 1) * sizeof(Uint32))) == NULL) {
		free(sax);
		return (-1);
	}


	sp = csp = (Uint8 *) src->pixels;
	dp = (Uint8 *) dst->pixels;
	dgap = dst->pitch - dst->w;

	if (flipx) csp += (src->w-1);
	if (flipy) csp  = ( (Uint8*)csp + src->pitch*(src->h-1) );


	csx = 0;
	csax = sax;
	for (x = 0; x < dst->w; x++) {
		csx += src->w;
		*csax = 0;
		while (csx >= dst->w) {
			csx -= dst->w;
			(*csax)++;
		}
		(*csax) = (*csax) * (flipx ? -1 : 1);
		csax++;
	}
	csy = 0;
	csay = say;
	for (y = 0; y < dst->h; y++) {
		csy += src->h;
		*csay = 0;
		while (csy >= dst->h) {
			csy -= dst->h;
			(*csay)++;
		}
		(*csay) = (*csay) * (flipy ? -1 : 1);
		csay++;
	}


	csay = say;
	for (y = 0; y < dst->h; y++) {
		csax = sax;
		sp = csp;
		for (x = 0; x < dst->w; x++) {

			*dp = *sp;

			sp += (*csax);
			csax++;

			dp++;
		}

		csp += ((*csay) * src->pitch);
		csay++;


		dp += dgap;
	}


	free(sax);
	free(say);

	return (0);
}


void _transformSurfaceRGBA(SDL_Surface * src, SDL_Surface * dst, int cx, int cy, int isin, int icos, int flipx, int flipy, int smooth)
{
	int x, y, t1, t2, dx, dy, xd, yd, sdx, sdy, ax, ay, ex, ey, sw, sh;
	tColorRGBA c00, c01, c10, c11, cswap;
	tColorRGBA *pc, *sp;
	int gap;


	xd = ((src->w - dst->w) << 15);
	yd = ((src->h - dst->h) << 15);
	ax = (cx << 16) - (icos * cx);
	ay = (cy << 16) - (isin * cx);
	sw = src->w - 1;
	sh = src->h - 1;
	pc = (tColorRGBA*) dst->pixels;
	gap = dst->pitch - dst->w * 4;


	if (smooth) {
		for (y = 0; y < dst->h; y++) {
			dy = cy - y;
			sdx = (ax + (isin * dy)) + xd;
			sdy = (ay - (icos * dy)) + yd;
			for (x = 0; x < dst->w; x++) {
				dx = (sdx >> 16);
				dy = (sdy >> 16);
				if (flipx) dx = sw - dx;
				if (flipy) dy = sh - dy;
				if ((dx > -1) && (dy > -1) && (dx < (src->w-1)) && (dy < (src->h-1))) {
					sp = (tColorRGBA *)src->pixels;;
					sp += ((src->pitch/4) * dy);
					sp += dx;
					c00 = *sp;
					sp += 1;
					c01 = *sp;
					sp += (src->pitch/4);
					c11 = *sp;
					sp -= 1;
					c10 = *sp;
					if (flipx) {
						cswap = c00; c00=c01; c01=cswap;
						cswap = c10; c10=c11; c11=cswap;
					}
					if (flipy) {
						cswap = c00; c00=c10; c10=cswap;
						cswap = c01; c01=c11; c11=cswap;
					}

					ex = (sdx & 0xffff);
					ey = (sdy & 0xffff);
					t1 = ((((c01.r - c00.r) * ex) >> 16) + c00.r) & 0xff;
					t2 = ((((c11.r - c10.r) * ex) >> 16) + c10.r) & 0xff;
					pc->r = (((t2 - t1) * ey) >> 16) + t1;
					t1 = ((((c01.g - c00.g) * ex) >> 16) + c00.g) & 0xff;
					t2 = ((((c11.g - c10.g) * ex) >> 16) + c10.g) & 0xff;
					pc->g = (((t2 - t1) * ey) >> 16) + t1;
					t1 = ((((c01.b - c00.b) * ex) >> 16) + c00.b) & 0xff;
					t2 = ((((c11.b - c10.b) * ex) >> 16) + c10.b) & 0xff;
					pc->b = (((t2 - t1) * ey) >> 16) + t1;
					t1 = ((((c01.a - c00.a) * ex) >> 16) + c00.a) & 0xff;
					t2 = ((((c11.a - c10.a) * ex) >> 16) + c10.a) & 0xff;
					pc->a = (((t2 - t1) * ey) >> 16) + t1;
				}
				sdx += icos;
				sdy += isin;
				pc++;
			}
			pc = (tColorRGBA *) ((Uint8 *) pc + gap);
		}
	} else {
		for (y = 0; y < dst->h; y++) {
			dy = cy - y;
			sdx = (ax + (isin * dy)) + xd;
			sdy = (ay - (icos * dy)) + yd;
			for (x = 0; x < dst->w; x++) {
				dx = (short) (sdx >> 16);
				dy = (short) (sdy >> 16);
				if (flipx) dx = (src->w-1)-dx;
				if (flipy) dy = (src->h-1)-dy;
				if ((dx >= 0) && (dy >= 0) && (dx < src->w) && (dy < src->h)) {
					sp = (tColorRGBA *) ((Uint8 *) src->pixels + src->pitch * dy);
					sp += dx;
					*pc = *sp;
				}
				sdx += icos;
				sdy += isin;
				pc++;
			}
			pc = (tColorRGBA *) ((Uint8 *) pc + gap);
		}
	}
}


void transformSurfaceY(SDL_Surface * src, SDL_Surface * dst, int cx, int cy, int isin, int icos, int flipx, int flipy)
{
	int x, y, dx, dy, xd, yd, sdx, sdy, ax, ay;
	tColorY *pc, *sp;
	int gap;


	xd = ((src->w - dst->w) << 15);
	yd = ((src->h - dst->h) << 15);
	ax = (cx << 16) - (icos * cx);
	ay = (cy << 16) - (isin * cx);
	pc = (tColorY*) dst->pixels;
	gap = dst->pitch - dst->w;

	memset(pc, (int)(_colorkey(src) & 0xff), dst->pitch * dst->h);

	for (y = 0; y < dst->h; y++) {
		dy = cy - y;
		sdx = (ax + (isin * dy)) + xd;
		sdy = (ay - (icos * dy)) + yd;
		for (x = 0; x < dst->w; x++) {
			dx = (short) (sdx >> 16);
			dy = (short) (sdy >> 16);
			if (flipx) dx = (src->w-1)-dx;
			if (flipy) dy = (src->h-1)-dy;
			if ((dx >= 0) && (dy >= 0) && (dx < src->w) && (dy < src->h)) {
				sp = (tColorY *) (src->pixels);
				sp += (src->pitch * dy + dx);
				*pc = *sp;
			}
			sdx += icos;
			sdy += isin;
			pc++;
		}
		pc += gap;
	}
}


SDL_Surface* rotateSurface90Degrees(SDL_Surface* src, int numClockwiseTurns)
{
	int row, col, newWidth, newHeight;
	int bpp, bpr;
	SDL_Surface* dst;
	Uint8* srcBuf;
	Uint8* dstBuf;
	int normalizedClockwiseTurns;


	if (!src ||
	    !src->format) {
		SDL_SetError("NULL source surface or source surface format");
	    return NULL;
	}

	if ((src->format->BitsPerPixel % 8) != 0) {
		SDL_SetError("Invalid source surface bit depth");
	    return NULL;
	}


	normalizedClockwiseTurns = (numClockwiseTurns % 4);
	if (normalizedClockwiseTurns < 0) {
		normalizedClockwiseTurns += 4;
	}


	if (normalizedClockwiseTurns % 2) {
		newWidth = src->h;
		newHeight = src->w;
	} else {
		newWidth = src->w;
		newHeight = src->h;
	}

	dst = SDL_CreateRGBSurface( src->flags, newWidth, newHeight, src->format->BitsPerPixel,
		src->format->Rmask,
		src->format->Gmask,
		src->format->Bmask,
		src->format->Amask);
	if(!dst) {
		SDL_SetError("Could not create destination surface");
		return NULL;
	}

	if (SDL_MUSTLOCK(src)) {
		SDL_LockSurface(src);
	}
	if (SDL_MUSTLOCK(dst)) {
		SDL_LockSurface(dst);
	}


	bpp = src->format->BitsPerPixel / 8;

	switch(normalizedClockwiseTurns) {
	case 0:
		{


			if (src->pitch == dst->pitch) {

				memcpy(dst->pixels, src->pixels, (src->h * src->pitch));
			}
			else
			{

				srcBuf = (Uint8*)(src->pixels);
				dstBuf = (Uint8*)(dst->pixels);
				bpr = src->w * bpp;
				for (row = 0; row < src->h; row++) {
					memcpy(dstBuf, srcBuf, bpr);
					srcBuf += src->pitch;
					dstBuf += dst->pitch;
				}
			}
		}
		break;


	case 1:
		{
			for (row = 0; row < src->h; ++row) {
				srcBuf = (Uint8*)(src->pixels) + (row * src->pitch);
				dstBuf = (Uint8*)(dst->pixels) + (dst->w - row - 1) * bpp;
				for (col = 0; col < src->w; ++col) {
					memcpy (dstBuf, srcBuf, bpp);
					srcBuf += bpp;
					dstBuf += dst->pitch;
				}
			}
		}
		break;

	case 2:
		{
			for (row = 0; row < src->h; ++row) {
				srcBuf = (Uint8*)(src->pixels) + (row * src->pitch);
				dstBuf = (Uint8*)(dst->pixels) + ((dst->h - row - 1) * dst->pitch) + (dst->w - 1) * bpp;
				for (col = 0; col < src->w; ++col) {
					memcpy (dstBuf, srcBuf, bpp);
					srcBuf += bpp;
					dstBuf -= bpp;
				}
			}
		}
		break;

	case 3:
		{
			for (row = 0; row < src->h; ++row) {
				srcBuf = (Uint8*)(src->pixels) + (row * src->pitch);
				dstBuf = (Uint8*)(dst->pixels) + (row * bpp) + ((dst->h - 1) * dst->pitch);
				for (col = 0; col < src->w; ++col) {
					memcpy (dstBuf, srcBuf, bpp);
					srcBuf += bpp;
					dstBuf -= dst->pitch;
				}
			}
		}
		break;
	}


	if (SDL_MUSTLOCK(src)) {
		SDL_UnlockSurface(src);
	}
	if (SDL_MUSTLOCK(dst)) {
		SDL_UnlockSurface(dst);
	}

	return dst;
}



void _rotozoomSurfaceSizeTrig(int width, int height, double angle, double zoomx, double zoomy,
	int *dstwidth, int *dstheight,
	double *canglezoom, double *sanglezoom)
{
	double x, y, cx, cy, sx, sy;
	double radangle;
	int dstwidthhalf, dstheighthalf;


	radangle = angle * (M_PI / 180.0);
	*sanglezoom = sin(radangle);
	*canglezoom = cos(radangle);
	*sanglezoom *= zoomx;
	*canglezoom *= zoomy;
	x = (double)(width / 2);
	y = (double)(height / 2);
	cx = *canglezoom * x;
	cy = *canglezoom * y;
	sx = *sanglezoom * x;
	sy = *sanglezoom * y;

	dstwidthhalf = MAX((int)
		ceil(MAX(MAX(MAX(fabs(cx + sy), fabs(cx - sy)), fabs(-cx + sy)), fabs(-cx - sy))), 1);
	dstheighthalf = MAX((int)
		ceil(MAX(MAX(MAX(fabs(sx + cy), fabs(sx - cy)), fabs(-sx + cy)), fabs(-sx - cy))), 1);
	*dstwidth = 2 * dstwidthhalf;
	*dstheight = 2 * dstheighthalf;
}


void rotozoomSurfaceSizeXY(int width, int height, double angle, double zoomx, double zoomy, int *dstwidth, int *dstheight)
{
	double dummy_sanglezoom, dummy_canglezoom;

	_rotozoomSurfaceSizeTrig(width, height, angle, zoomx, zoomy, dstwidth, dstheight, &dummy_sanglezoom, &dummy_canglezoom);
}


void rotozoomSurfaceSize(int width, int height, double angle, double zoom, int *dstwidth, int *dstheight)
{
	double dummy_sanglezoom, dummy_canglezoom;

	_rotozoomSurfaceSizeTrig(width, height, angle, zoom, zoom, dstwidth, dstheight, &dummy_sanglezoom, &dummy_canglezoom);
}


SDL_Surface *rotozoomSurface(SDL_Surface * src, double angle, double zoom, int smooth)
{
	return rotozoomSurfaceXY(src, angle, zoom, zoom, smooth);
}


SDL_Surface *rotozoomSurfaceXY(SDL_Surface * src, double angle, double zoomx, double zoomy, int smooth)
{
	SDL_Surface *rz_src;
	SDL_Surface *rz_dst;
	double zoominv;
	double sanglezoom, canglezoom, sanglezoominv, canglezoominv;
	int dstwidthhalf, dstwidth, dstheighthalf, dstheight;
	int is32bit;
	int i, src_converted;
	int flipx,flipy;


	if (src == NULL) {
		return (NULL);
	}


	is32bit = (src->format->BitsPerPixel == 32);
	if ((is32bit) || (src->format->BitsPerPixel == 8)) {

		rz_src = src;
		src_converted = 0;
	} else {

		rz_src =
			SDL_CreateRGBSurface(SDL_SWSURFACE, src->w, src->h, 32,
#if SDL_BYTEORDER == SDL_LIL_ENDIAN
			0x000000ff, 0x0000ff00, 0x00ff0000, 0xff000000
#else
			0xff000000,  0x00ff0000, 0x0000ff00, 0x000000ff
#endif
			);

		SDL_BlitSurface(src, NULL, rz_src, NULL);

		src_converted = 1;
		is32bit = 1;
	}


	flipx = (zoomx<0.0);
	if (flipx) zoomx=-zoomx;
	flipy = (zoomy<0.0);
	if (flipy) zoomy=-zoomy;
	if (zoomx < VALUE_LIMIT) zoomx = VALUE_LIMIT;
	if (zoomy < VALUE_LIMIT) zoomy = VALUE_LIMIT;
	zoominv = 65536.0 / (zoomx * zoomx);


	if (fabs(angle) > VALUE_LIMIT) {





		_rotozoomSurfaceSizeTrig(rz_src->w, rz_src->h, angle, zoomx, zoomy, &dstwidth, &dstheight, &canglezoom, &sanglezoom);


		sanglezoominv = sanglezoom;
		canglezoominv = canglezoom;
		sanglezoominv *= zoominv;
		canglezoominv *= zoominv;


		dstwidthhalf = dstwidth / 2;
		dstheighthalf = dstheight / 2;


		rz_dst = NULL;
		if (is32bit) {

			rz_dst =
				SDL_CreateRGBSurface(SDL_SWSURFACE, dstwidth, dstheight + GUARD_ROWS, 32,
				rz_src->format->Rmask, rz_src->format->Gmask,
				rz_src->format->Bmask, rz_src->format->Amask);
		} else {

			rz_dst = SDL_CreateRGBSurface(SDL_SWSURFACE, dstwidth, dstheight + GUARD_ROWS, 8, 0, 0, 0, 0);
		}


		if (rz_dst == NULL)
			return NULL;


		rz_dst->h = dstheight;


		if (SDL_MUSTLOCK(rz_src)) {
			SDL_LockSurface(rz_src);
		}


		if (is32bit) {

			_transformSurfaceRGBA(rz_src, rz_dst, dstwidthhalf, dstheighthalf,
				(int) (sanglezoominv), (int) (canglezoominv),
				flipx, flipy,
				smooth);
		} else {

			for (i = 0; i < rz_src->format->palette->ncolors; i++) {
				rz_dst->format->palette->colors[i] = rz_src->format->palette->colors[i];
			}
			rz_dst->format->palette->ncolors = rz_src->format->palette->ncolors;

			transformSurfaceY(rz_src, rz_dst, dstwidthhalf, dstheighthalf,
				(int) (sanglezoominv), (int) (canglezoominv),
				flipx, flipy);
		}

		if (SDL_MUSTLOCK(rz_src)) {
			SDL_UnlockSurface(rz_src);
		}

	} else {





		zoomSurfaceSize(rz_src->w, rz_src->h, zoomx, zoomy, &dstwidth, &dstheight);


		rz_dst = NULL;
		if (is32bit) {

			rz_dst =
				SDL_CreateRGBSurface(SDL_SWSURFACE, dstwidth, dstheight + GUARD_ROWS, 32,
				rz_src->format->Rmask, rz_src->format->Gmask,
				rz_src->format->Bmask, rz_src->format->Amask);
		} else {

			rz_dst = SDL_CreateRGBSurface(SDL_SWSURFACE, dstwidth, dstheight + GUARD_ROWS, 8, 0, 0, 0, 0);
		}


		if (rz_dst == NULL)
			return NULL;


		rz_dst->h = dstheight;


		if (SDL_MUSTLOCK(rz_src)) {
			SDL_LockSurface(rz_src);
		}


		if (is32bit) {

			_zoomSurfaceRGBA(rz_src, rz_dst, flipx, flipy, smooth);

		} else {

			for (i = 0; i < rz_src->format->palette->ncolors; i++) {
				rz_dst->format->palette->colors[i] = rz_src->format->palette->colors[i];
			}
			rz_dst->format->palette->ncolors = rz_src->format->palette->ncolors;


			_zoomSurfaceY(rz_src, rz_dst, flipx, flipy);
		}


		if (SDL_MUSTLOCK(rz_src)) {
			SDL_UnlockSurface(rz_src);
		}
	}


	if (src_converted) {
		SDL_FreeSurface(rz_src);
	}


	return (rz_dst);
}


void zoomSurfaceSize(int width, int height, double zoomx, double zoomy, int *dstwidth, int *dstheight)
{

	int flipx, flipy;
	flipx = (zoomx<0.0);
	if (flipx) zoomx = -zoomx;
	flipy = (zoomy<0.0);
	if (flipy) zoomy = -zoomy;


	if (zoomx < VALUE_LIMIT) {
		zoomx = VALUE_LIMIT;
	}
	if (zoomy < VALUE_LIMIT) {
		zoomy = VALUE_LIMIT;
	}


	*dstwidth = (int) floor(((double) width * zoomx) + 0.5);
	*dstheight = (int) floor(((double) height * zoomy) + 0.5);
	if (*dstwidth < 1) {
		*dstwidth = 1;
	}
	if (*dstheight < 1) {
		*dstheight = 1;
	}
}


SDL_Surface *zoomSurface(SDL_Surface * src, double zoomx, double zoomy, int smooth)
{
	SDL_Surface *rz_src;
	SDL_Surface *rz_dst;
	int dstwidth, dstheight;
	int is32bit;
	int i, src_converted;
	int flipx, flipy;


	if (src == NULL)
		return (NULL);


	is32bit = (src->format->BitsPerPixel == 32);
	if ((is32bit) || (src->format->BitsPerPixel == 8)) {

		rz_src = src;
		src_converted = 0;
	} else {

		rz_src =
			SDL_CreateRGBSurface(SDL_SWSURFACE, src->w, src->h, 32,
#if SDL_BYTEORDER == SDL_LIL_ENDIAN
			0x000000ff, 0x0000ff00, 0x00ff0000, 0xff000000
#else
			0xff000000,  0x00ff0000, 0x0000ff00, 0x000000ff
#endif
			);
		if (rz_src == NULL) {
			return NULL;
		}
		SDL_BlitSurface(src, NULL, rz_src, NULL);
		src_converted = 1;
		is32bit = 1;
	}

	flipx = (zoomx<0.0);
	if (flipx) zoomx = -zoomx;
	flipy = (zoomy<0.0);
	if (flipy) zoomy = -zoomy;


	zoomSurfaceSize(rz_src->w, rz_src->h, zoomx, zoomy, &dstwidth, &dstheight);


	rz_dst = NULL;
	if (is32bit) {

		rz_dst =
			SDL_CreateRGBSurface(SDL_SWSURFACE, dstwidth, dstheight + GUARD_ROWS, 32,
			rz_src->format->Rmask, rz_src->format->Gmask,
			rz_src->format->Bmask, rz_src->format->Amask);
	} else {

		rz_dst = SDL_CreateRGBSurface(SDL_SWSURFACE, dstwidth, dstheight + GUARD_ROWS, 8, 0, 0, 0, 0);
	}


	if (rz_dst == NULL) {

		if (src_converted) {
			SDL_FreeSurface(rz_src);
		}
		return NULL;
	}


	rz_dst->h = dstheight;


	if (SDL_MUSTLOCK(rz_src)) {
		SDL_LockSurface(rz_src);
	}


	if (is32bit) {

		_zoomSurfaceRGBA(rz_src, rz_dst, flipx, flipy, smooth);
	} else {

		for (i = 0; i < rz_src->format->palette->ncolors; i++) {
			rz_dst->format->palette->colors[i] = rz_src->format->palette->colors[i];
		}
		rz_dst->format->palette->ncolors = rz_src->format->palette->ncolors;

		_zoomSurfaceY(rz_src, rz_dst, flipx, flipy);
	}

	if (SDL_MUSTLOCK(rz_src)) {
		SDL_UnlockSurface(rz_src);
	}


	if (src_converted) {
		SDL_FreeSurface(rz_src);
	}


	return (rz_dst);
}



SDL_Surface *shrinkSurface(SDL_Surface *src, int factorx, int factory)
{
	int result;
	SDL_Surface *rz_src;
	SDL_Surface *rz_dst = NULL;
	int dstwidth, dstheight;
	int is32bit;
	int i, src_converted;
	int haveError = 0;


	if (src == NULL) {
		return (NULL);
	}


	is32bit = (src->format->BitsPerPixel == 32);
	if ((is32bit) || (src->format->BitsPerPixel == 8)) {

		rz_src = src;
		src_converted = 0;
	} else {

		rz_src = SDL_CreateRGBSurface(SDL_SWSURFACE, src->w, src->h, 32,
#if SDL_BYTEORDER == SDL_LIL_ENDIAN
			0x000000ff, 0x0000ff00, 0x00ff0000, 0xff000000
#else
			0xff000000,  0x00ff0000, 0x0000ff00, 0x000000ff
#endif
			);
		if (rz_src==NULL) {
			haveError = 1;
			goto exitShrinkSurface;
		}

		SDL_BlitSurface(src, NULL, rz_src, NULL);
		src_converted = 1;
		is32bit = 1;
	}


	if (SDL_MUSTLOCK(rz_src)) {
		if (SDL_LockSurface(rz_src) < 0) {
			haveError = 1;
			goto exitShrinkSurface;
		}
	}


	dstwidth=rz_src->w/factorx;
	while (dstwidth*factorx>rz_src->w) { dstwidth--; }
	dstheight=rz_src->h/factory;
	while (dstheight*factory>rz_src->h) { dstheight--; }


	if (is32bit==1) {

		rz_dst =
			SDL_CreateRGBSurface(SDL_SWSURFACE, dstwidth, dstheight + GUARD_ROWS, 32,
			rz_src->format->Rmask, rz_src->format->Gmask,
			rz_src->format->Bmask, rz_src->format->Amask);
	} else {

		rz_dst = SDL_CreateRGBSurface(SDL_SWSURFACE, dstwidth, dstheight + GUARD_ROWS, 8, 0, 0, 0, 0);
	}


	if (rz_dst == NULL) {
		haveError = 1;
		goto exitShrinkSurface;
	}


	rz_dst->h = dstheight;


	if (is32bit==1) {

		result = _shrinkSurfaceRGBA(rz_src, rz_dst, factorx, factory);
		if ((result!=0) || (rz_dst==NULL)) {
			haveError = 1;
			goto exitShrinkSurface;
		}
	} else {

		for (i = 0; i < rz_src->format->palette->ncolors; i++) {
			rz_dst->format->palette->colors[i] = rz_src->format->palette->colors[i];
		}
		rz_dst->format->palette->ncolors = rz_src->format->palette->ncolors;

		result = _shrinkSurfaceY(rz_src, rz_dst, factorx, factory);
		if (result!=0) {
			haveError = 1;
			goto exitShrinkSurface;
		}
	}

exitShrinkSurface:
	if (rz_src!=NULL) {

		if (SDL_MUSTLOCK(rz_src)) {
			SDL_UnlockSurface(rz_src);
		}


		if (src_converted==1) {
			SDL_FreeSurface(rz_src);
		}
	}


	if (haveError==1) {
		if (rz_dst!=NULL) {
			SDL_FreeSurface(rz_dst);
		}
		rz_dst=NULL;
	}


	return (rz_dst);
}
*/
import (
	"C"
)
