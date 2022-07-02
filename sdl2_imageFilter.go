package gfx

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "SDL2/SDL.h"


#ifdef __GNUC__
#  ifdef USE_MMX
#    include <mmintrin.h>
#  endif
#  include <SDL2/SDL_cpuinfo.h>
#endif

#include "SDL2/SDL2_imageFilter.h"


#define SWAP_32(x) (((x) >> 24) | (((x) & 0x00ff0000) >> 8)  | (((x) & 0x0000ff00) << 8)  | ((x) << 24))




static int SDL_imageFilterUseMMX = 1;


#if defined(__GNUC__)
#define GCC__
#endif


int SDL_imageFilterMMXdetect(void)
{

if (SDL_imageFilterUseMMX == 0) {
return (0);
}

return SDL_HasMMX();
}


void SDL_imageFilterMMXoff()
{
SDL_imageFilterUseMMX = 0;
}


void SDL_imageFilterMMXon()
{
SDL_imageFilterUseMMX = 1;
}




static int SDL_imageFilterAddMMX(unsigned char *Src1, unsigned char *Src2, unsigned char *Dest, unsigned int SrcLength)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
pusha
mov eax, Src1
mov ebx, Src2
mov edi, Dest
mov ecx, SrcLength
shr ecx, 3
align 16
L1010:
movq mm1, [eax]
paddusb mm1, [ebx]
movq [edi], mm1
add eax, 8
add ebx, 8
add edi, 8
dec ecx
jnz L1010
emms
popa
}
#else

__m64 *mSrc1 = (__m64*)Src1;
__m64 *mSrc2 = (__m64*)Src2;
__m64 *mDest = (__m64*)Dest;
int i;
for (i = 0; i < SrcLength/8; i++) {
*mDest = _m_paddusb(*mSrc1, *mSrc2);
mSrc1++;
mSrc2++;
mDest++;
}
_m_empty();
#endif
return (0);
#else
return (-1);
#endif
}


int SDL_imageFilterAdd(unsigned char *Src1, unsigned char *Src2, unsigned char *Dest, unsigned int length)
{
unsigned int i, istart;
unsigned char *cursrc1, *cursrc2, *curdst;
int result;


if ((Src1 == NULL) || (Src2 == NULL) || (Dest == NULL))
return(-1);
if (length == 0)
return(0);

if ((SDL_imageFilterMMXdetect()) && (length > 7)) {


SDL_imageFilterAddMMX(Src1, Src2, Dest, length);


if ((length & 7) > 0) {

istart = length & 0xfffffff8;
cursrc1 = &Src1[istart];
cursrc2 = &Src2[istart];
curdst = &Dest[istart];
} else {

return (0);
}
} else {

istart = 0;
cursrc1 = Src1;
cursrc2 = Src2;
curdst = Dest;
}


for (i = istart; i < length; i++) {
result = (int) *cursrc1 + (int) *cursrc2;
if (result > 255)
result = 255;
*curdst = (unsigned char) result;

cursrc1++;
cursrc2++;
curdst++;
}

return (0);
}


static int SDL_imageFilterMeanMMX(unsigned char *Src1, unsigned char *Src2, unsigned char *Dest, unsigned int SrcLength,
unsigned char *Mask)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
pusha
mov edx, Mask
movq mm0, [edx]
mov eax, Src1
mov ebx, Src2
mov edi, Dest
mov ecx, SrcLength
shr ecx, 3
align 16
L21011:
movq mm1,  [eax]
movq mm2,  [ebx]

psrlw mm1, 1
psrlw mm2, 1
pand mm1, mm0

pand mm2, mm0

paddusb mm1,  mm2
movq [edi],  mm1
add eax,  8
add ebx,  8
add edi,  8
dec ecx
jnz L21011
emms
popa
}
#else

__m64 *mSrc1 = (__m64*)Src1;
__m64 *mSrc2 = (__m64*)Src2;
__m64 *mDest = (__m64*)Dest;
__m64 *mMask = (__m64*)Mask;
int i;
for (i = 0; i < SrcLength/8; i++) {
__m64 mm1 = *mSrc1,
mm2 = *mSrc2;
mm1 = _m_psrlwi(mm1, 1);
mm2 = _m_psrlwi(mm2, 1);
mm1 = _m_pand(mm1, *mMask);
mm2 = _m_pand(mm2, *mMask);
*mDest = _m_paddusb(mm1, mm2);
mSrc1++;
mSrc2++;
mDest++;
}
_m_empty();
#endif
return (0);
#else
return (-1);
#endif
}


int SDL_imageFilterMean(unsigned char *Src1, unsigned char *Src2, unsigned char *Dest, unsigned int length)
{
static unsigned char Mask[8] = { 0x7F, 0x7F, 0x7F, 0x7F, 0x7F, 0x7F, 0x7F, 0x7F };
unsigned int i, istart;
unsigned char *cursrc1, *cursrc2, *curdst;
int result;


if ((Src1 == NULL) || (Src2 == NULL) || (Dest == NULL))
return(-1);
if (length == 0)
return(0);

if ((SDL_imageFilterMMXdetect()) && (length > 7)) {

SDL_imageFilterMeanMMX(Src1, Src2, Dest, length, Mask);


if ((length & 7) > 0) {

istart = length & 0xfffffff8;
cursrc1 = &Src1[istart];
cursrc2 = &Src2[istart];
curdst = &Dest[istart];
} else {

return (0);
}
} else {

istart = 0;
cursrc1 = Src1;
cursrc2 = Src2;
curdst = Dest;
}


for (i = istart; i < length; i++) {
result = (int) *cursrc1 / 2 + (int) *cursrc2 / 2;
*curdst = (unsigned char) result;

cursrc1++;
cursrc2++;
curdst++;
}

return (0);
}


static int SDL_imageFilterSubMMX(unsigned char *Src1, unsigned char *Src2, unsigned char *Dest, unsigned int SrcLength)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
pusha
mov eax,  Src1
mov ebx,  Src2
mov edi,  Dest
mov ecx,  SrcLength
shr ecx,  3
align 16
L1012:
movq mm1,  [eax]
psubusb mm1,  [ebx]
movq [edi],  mm1
add eax, 8
add ebx, 8
add edi, 8
dec ecx
jnz L1012
emms
popa
}
#else

__m64 *mSrc1 = (__m64*)Src1;
__m64 *mSrc2 = (__m64*)Src2;
__m64 *mDest = (__m64*)Dest;
int i;
for (i = 0; i < SrcLength/8; i++) {
*mDest = _m_psubusb(*mSrc1, *mSrc2);
mSrc1++;
mSrc2++;
mDest++;
}
_m_empty();
#endif
return (0);
#else
return (-1);
#endif
}


int SDL_imageFilterSub(unsigned char *Src1, unsigned char *Src2, unsigned char *Dest, unsigned int length)
{
unsigned int i, istart;
unsigned char *cursrc1, *cursrc2, *curdst;
int result;


if ((Src1 == NULL) || (Src2 == NULL) || (Dest == NULL))
return(-1);
if (length == 0)
return(0);

if ((SDL_imageFilterMMXdetect()) && (length > 7)) {

SDL_imageFilterSubMMX(Src1, Src2, Dest, length);


if ((length & 7) > 0) {

istart = length & 0xfffffff8;
cursrc1 = &Src1[istart];
cursrc2 = &Src2[istart];
curdst = &Dest[istart];
} else {

return (0);
}
} else {

istart = 0;
cursrc1 = Src1;
cursrc2 = Src2;
curdst = Dest;
}


for (i = istart; i < length; i++) {
result = (int) *cursrc1 - (int) *cursrc2;
if (result < 0)
result = 0;
*curdst = (unsigned char) result;

cursrc1++;
cursrc2++;
curdst++;
}

return (0);
}


static int SDL_imageFilterAbsDiffMMX(unsigned char *Src1, unsigned char *Src2, unsigned char *Dest, unsigned int SrcLength)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
pusha
mov eax, Src1
mov ebx, Src2
mov edi, Dest
mov ecx, SrcLength
shr ecx,  3
align 16
L1013:
movq mm1,  [eax]
movq mm2,  [ebx]
psubusb mm1,  [ebx]
psubusb mm2,  [eax]
por mm1,  mm2
movq [edi],  mm1
add eax, 8
add ebx, 8
add edi, 8
dec ecx
jnz L1013
emms
popa
}
#else

__m64 *mSrc1 = (__m64*)Src1;
__m64 *mSrc2 = (__m64*)Src2;
__m64 *mDest = (__m64*)Dest;
int i;
for (i = 0; i < SrcLength/8; i++) {
__m64 mm1 = _m_psubusb(*mSrc2, *mSrc1);
__m64 mm2 = _m_psubusb(*mSrc1, *mSrc2);
*mDest = _m_por(mm1, mm2);
mSrc1++;
mSrc2++;
mDest++;
}
_m_empty();
#endif
return (0);
#else
return (-1);
#endif
}


int SDL_imageFilterAbsDiff(unsigned char *Src1, unsigned char *Src2, unsigned char *Dest, unsigned int length)
{
unsigned int i, istart;
unsigned char *cursrc1, *cursrc2, *curdst;
int result;


if ((Src1 == NULL) || (Src2 == NULL) || (Dest == NULL))
return(-1);
if (length == 0)
return(0);

if ((SDL_imageFilterMMXdetect()) && (length > 7)) {

SDL_imageFilterAbsDiffMMX(Src1, Src2, Dest, length);


if ((length & 7) > 0) {

istart = length & 0xfffffff8;
cursrc1 = &Src1[istart];
cursrc2 = &Src2[istart];
curdst = &Dest[istart];
} else {

return (0);
}
} else {

istart = 0;
cursrc1 = Src1;
cursrc2 = Src2;
curdst = Dest;
}


for (i = istart; i < length; i++) {
result = abs((int) *cursrc1 - (int) *cursrc2);
*curdst = (unsigned char) result;

cursrc1++;
cursrc2++;
curdst++;
}

return (0);
}


static int SDL_imageFilterMultMMX(unsigned char *Src1, unsigned char *Src2, unsigned char *Dest, unsigned int SrcLength)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
pusha
mov eax, Src1
mov ebx, Src2
mov edi, Dest
mov ecx, SrcLength
shr ecx, 3
pxor mm0, mm0
align 16
L1014:
movq mm1, [eax]
movq mm3, [ebx]
movq mm2, mm1
movq mm4, mm3
punpcklbw mm1, mm0
punpckhbw mm2, mm0
punpcklbw mm3, mm0
punpckhbw mm4, mm0
pmullw mm1, mm3
pmullw mm2, mm4

movq mm5, mm1
movq mm6, mm2
psraw mm5, 15
psraw mm6, 15
pxor mm1, mm5
pxor mm2, mm6
psubsw mm1, mm5
psubsw mm2, mm6
packuswb mm1, mm2
movq [edi], mm1
add eax, 8
add ebx, 8
add edi, 8
dec ecx
jnz L1014
emms
popa
}
#else










































__m64 *mSrc1 = (__m64*)Src1;
__m64 *mSrc2 = (__m64*)Src2;
__m64 *mDest = (__m64*)Dest;
__m64 mm0 = _m_from_int(0);
int i;
for (i = 0; i < SrcLength/8; i++) {
__m64 mm1, mm2, mm3, mm4, mm5, mm6;
mm1 = _m_punpcklbw(*mSrc1, mm0);
mm2 = _m_punpckhbw(*mSrc1, mm0);
mm3 = _m_punpcklbw(*mSrc2, mm0);
mm4 = _m_punpckhbw(*mSrc2, mm0);
mm1 = _m_pmullw(mm1, mm3);
mm2 = _m_pmullw(mm2, mm4);
mm5 = _m_psrawi(mm1, 15);
mm6 = _m_psrawi(mm2, 15);
mm1 = _m_pxor(mm1, mm5);
mm2 = _m_pxor(mm2, mm6);
mm1 = _m_psubsw(mm1, mm5);
mm2 = _m_psubsw(mm2, mm6);
*mDest = _m_packuswb(mm1, mm2);
mSrc1++;
mSrc2++;
mDest++;
}
_m_empty();
#endif
return (0);
#else
return (-1);
#endif
}


int SDL_imageFilterMult(unsigned char *Src1, unsigned char *Src2, unsigned char *Dest, unsigned int length)
{
unsigned int i, istart;
unsigned char *cursrc1, *cursrc2, *curdst;
int result;


if ((Src1 == NULL) || (Src2 == NULL) || (Dest == NULL))
return(-1);
if (length == 0)
return(0);

if ((SDL_imageFilterMMXdetect()) && (length > 7)) {

SDL_imageFilterMultMMX(Src1, Src2, Dest, length);


if ((length & 7) > 0) {

istart = length & 0xfffffff8;
cursrc1 = &Src1[istart];
cursrc2 = &Src2[istart];
curdst = &Dest[istart];
} else {

return (0);
}
} else {

istart = 0;
cursrc1 = Src1;
cursrc2 = Src2;
curdst = Dest;
}


for (i = istart; i < length; i++) {



result = (int) *cursrc1 * (int) *cursrc2;
if (result > 255)
result = 255;
*curdst = (unsigned char) result;

cursrc1++;
cursrc2++;
curdst++;
}

return (0);
}


int SDL_imageFilterMultNorASM(unsigned char *Src1, unsigned char *Src2, unsigned char *Dest, unsigned int SrcLength)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
pusha
mov edx, Src1
mov esi, Src2
mov edi, Dest
mov ecx, SrcLength
align 16
L10141:
mov al, [edx]
mul [esi]
mov [edi], al
inc edx
inc esi
inc edi
dec ecx
jnz L10141
popa
}
#else


asm volatile (
".align 16       \n\t"
#  if defined(i386)
"1:mov  (%%edx), %%al \n\t"
"mulb (%%esi)       \n\t"
"mov %%al, (%%edi)  \n\t"
"inc %%edx \n\t"
"inc %%esi \n\t"
"inc %%edi \n\t"
"dec %%ecx      \n\t"
#  elif defined(__x86_64__)
"1:mov  (%%rdx), %%al \n\t"
"mulb (%%rsi)       \n\t"
"mov %%al, (%%rdi)  \n\t"
"inc %%rdx \n\t"
"inc %%rsi \n\t"
"inc %%rdi \n\t"
"dec %%rcx      \n\t"
#  endif
"jnz 1b         \n\t"
: "+d" (Src1),
"+S" (Src2),
"+c" (SrcLength),
"+D" (Dest)
:
: "memory", "rax"
);
#endif
return (0);
#else
return (-1);
#endif
}


int SDL_imageFilterMultNor(unsigned char *Src1, unsigned char *Src2, unsigned char *Dest, unsigned int length)
{
unsigned int i, istart;
unsigned char *cursrc1, *cursrc2, *curdst;


if ((Src1 == NULL) || (Src2 == NULL) || (Dest == NULL))
return(-1);
if (length == 0)
return(0);

if (SDL_imageFilterMMXdetect()) {
if (length > 0) {

SDL_imageFilterMultNorASM(Src1, Src2, Dest, length);


if ((length & 7) > 0) {

istart = length & 0xfffffff8;
cursrc1 = &Src1[istart];
cursrc2 = &Src2[istart];
curdst = &Dest[istart];
} else {

return (0);
}
} else {

return (0);
}
} else {

istart = 0;
cursrc1 = Src1;
cursrc2 = Src2;
curdst = Dest;
}


for (i = istart; i < length; i++) {
*curdst = (int)*cursrc1 * (int)*cursrc2;  // (int) for efficiency

cursrc1++;
cursrc2++;
curdst++;
}

return (0);
}


static int SDL_imageFilterMultDivby2MMX(unsigned char *Src1, unsigned char *Src2, unsigned char *Dest, unsigned int SrcLength)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
pusha
mov eax, Src1
mov ebx, Src2
mov edi, Dest
mov ecx,  SrcLength
shr ecx,  3
pxor mm0,  mm0
align 16
L1015:
movq mm1,  [eax]
movq mm3,  [ebx]
movq mm2,  mm1
movq mm4,  mm3
punpcklbw mm1,  mm0
punpckhbw mm2,  mm0
punpcklbw mm3,  mm0
punpckhbw mm4,  mm0
psrlw mm1,  1
psrlw mm2,  1
pmullw mm1,  mm3
pmullw mm2,  mm4
packuswb mm1,  mm2
movq [edi],  mm1
add eax,  8
add ebx,  8
add edi,  8
dec ecx
jnz L1015
emms
popa
}
#else

__m64 *mSrc1 = (__m64*)Src1;
__m64 *mSrc2 = (__m64*)Src2;
__m64 *mDest = (__m64*)Dest;
__m64 mm0 = _m_from_int(0);
int i;
for (i = 0; i < SrcLength/8; i++) {
__m64 mm1, mm2, mm3, mm4, mm5, mm6;
mm1 = _m_punpcklbw(*mSrc1, mm0);
mm2 = _m_punpckhbw(*mSrc1, mm0);
mm3 = _m_punpcklbw(*mSrc2, mm0);
mm4 = _m_punpckhbw(*mSrc2, mm0);
mm1 = _m_psrlwi(mm1, 1);
mm2 = _m_psrlwi(mm2, 1);
mm1 = _m_pmullw(mm1, mm3);
mm2 = _m_pmullw(mm2, mm4);
*mDest = _m_packuswb(mm1, mm2);
mSrc1++;
mSrc2++;
mDest++;
}
_m_empty();
#endif
return (0);
#else
return (-1);
#endif
}


int SDL_imageFilterMultDivby2(unsigned char *Src1, unsigned char *Src2, unsigned char *Dest, unsigned int length)
{
unsigned int i, istart;
unsigned char *cursrc1, *cursrc2, *curdst;
int result;


if ((Src1 == NULL) || (Src2 == NULL) || (Dest == NULL))
return(-1);
if (length == 0)
return(0);

if ((SDL_imageFilterMMXdetect()) && (length > 7)) {

SDL_imageFilterMultDivby2MMX(Src1, Src2, Dest, length);


if ((length & 7) > 0) {

istart = length & 0xfffffff8;
cursrc1 = &Src1[istart];
cursrc2 = &Src2[istart];
curdst = &Dest[istart];
} else {

return (0);
}
} else {

istart = 0;
cursrc1 = Src1;
cursrc2 = Src2;
curdst = Dest;
}


for (i = istart; i < length; i++) {
result = ((int) *cursrc1 / 2) * (int) *cursrc2;
if (result > 255)
result = 255;
*curdst = (unsigned char) result;

cursrc1++;
cursrc2++;
curdst++;
}

return (0);
}


static int SDL_imageFilterMultDivby4MMX(unsigned char *Src1, unsigned char *Src2, unsigned char *Dest, unsigned int SrcLength)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
pusha
mov eax, Src1
mov ebx, Src2
mov edi, Dest
mov ecx, SrcLength
shr ecx,  3
pxor mm0, mm0
align 16
L1016:
movq mm1, [eax]
movq mm3, [ebx]
movq mm2, mm1
movq mm4, mm3
punpcklbw mm1, mm0
punpckhbw mm2, mm0
punpcklbw mm3, mm0
punpckhbw mm4, mm0
psrlw mm1, 1
psrlw mm2, 1
psrlw mm3, 1
psrlw mm4, 1
pmullw mm1, mm3
pmullw mm2, mm4
packuswb mm1, mm2
movq [edi], mm1
add eax, 8
add ebx, 8
add edi,  8
dec ecx
jnz L1016
emms
popa
}
#else

__m64 *mSrc1 = (__m64*)Src1;
__m64 *mSrc2 = (__m64*)Src2;
__m64 *mDest = (__m64*)Dest;
__m64 mm0 = _m_from_int(0);
int i;
for (i = 0; i < SrcLength/8; i++) {
__m64 mm1, mm2, mm3, mm4, mm5, mm6;
mm1 = _m_punpcklbw(*mSrc1, mm0);
mm2 = _m_punpckhbw(*mSrc1, mm0);
mm3 = _m_punpcklbw(*mSrc2, mm0);
mm4 = _m_punpckhbw(*mSrc2, mm0);
mm1 = _m_psrlwi(mm1, 1);
mm2 = _m_psrlwi(mm2, 1);
mm3 = _m_psrlwi(mm3, 1);
mm4 = _m_psrlwi(mm4, 1);
mm1 = _m_pmullw(mm1, mm3);
mm2 = _m_pmullw(mm2, mm4);
*mDest = _m_packuswb(mm1, mm2);
mSrc1++;
mSrc2++;
mDest++;
}
_m_empty();
#endif
return (0);
#else
return (-1);
#endif
}


int SDL_imageFilterMultDivby4(unsigned char *Src1, unsigned char *Src2, unsigned char *Dest, unsigned int length)
{
unsigned int i, istart;
unsigned char *cursrc1, *cursrc2, *curdst;
int result;


if ((Src1 == NULL) || (Src2 == NULL) || (Dest == NULL))
return(-1);
if (length == 0)
return(0);

if ((SDL_imageFilterMMXdetect()) && (length > 7)) {

SDL_imageFilterMultDivby4MMX(Src1, Src2, Dest, length);


if ((length & 7) > 0) {

istart = length & 0xfffffff8;
cursrc1 = &Src1[istart];
cursrc2 = &Src2[istart];
curdst = &Dest[istart];
} else {

return (0);
}
} else {

istart = 0;
cursrc1 = Src1;
cursrc2 = Src2;
curdst = Dest;
}


for (i = istart; i < length; i++) {
result = ((int) *cursrc1 / 2) * ((int) *cursrc2 / 2);
if (result > 255)
result = 255;
*curdst = (unsigned char) result;

cursrc1++;
cursrc2++;
curdst++;
}

return (0);
}


static int SDL_imageFilterBitAndMMX(unsigned char *Src1, unsigned char *Src2, unsigned char *Dest, unsigned int SrcLength)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
pusha
mov eax, Src1
mov ebx, Src2
mov edi, Dest
mov ecx, SrcLength
shr ecx, 3
align 16
L1017:
movq mm1, [eax]
pand mm1, [ebx]
movq [edi], mm1
add eax, 8
add ebx, 8
add edi, 8
dec ecx
jnz L1017
emms
popa
}
#else























__m64 *mSrc1 = (__m64*)Src1;
__m64 *mSrc2 = (__m64*)Src2;
__m64 *mDest = (__m64*)Dest;
int i;
for (i = 0; i < SrcLength/8; i++) {
*mDest = _m_pand(*mSrc1, *mSrc2);
mSrc1++;
mSrc2++;
mDest++;
}
_m_empty();
#endif
return (0);
#else
return (-1);
#endif
}


int SDL_imageFilterBitAnd(unsigned char *Src1, unsigned char *Src2, unsigned char *Dest, unsigned int length)
{
unsigned int i, istart;
unsigned char *cursrc1, *cursrc2, *curdst;


if ((Src1 == NULL) || (Src2 == NULL) || (Dest == NULL))
return(-1);
if (length == 0)
return(0);

if ((SDL_imageFilterMMXdetect()>0) && (length>7)) {



SDL_imageFilterBitAndMMX(Src1, Src2, Dest, length);


if ((length & 7) > 0) {


istart = length & 0xfffffff8;
cursrc1 = &Src1[istart];
cursrc2 = &Src2[istart];
curdst = &Dest[istart];
} else {

return (0);
}
} else {

istart = 0;
cursrc1 = Src1;
cursrc2 = Src2;
curdst = Dest;
}


for (i = istart; i < length; i++) {
*curdst = (*cursrc1) & (*cursrc2);

cursrc1++;
cursrc2++;
curdst++;
}

return (0);
}


static int SDL_imageFilterBitOrMMX(unsigned char *Src1, unsigned char *Src2, unsigned char *Dest, unsigned int SrcLength)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
pusha
mov eax, Src1
mov ebx, Src2
mov edi, Dest
mov ecx, SrcLength
shr ecx,  3
align 16
L91017:
movq mm1, [eax]
por mm1, [ebx]
movq [edi], mm1
add eax, 8
add ebx, 8
add edi,  8
dec ecx
jnz L91017
emms
popa
}
#else

__m64 *mSrc1 = (__m64*)Src1;
__m64 *mSrc2 = (__m64*)Src2;
__m64 *mDest = (__m64*)Dest;
int i;
for (i = 0; i < SrcLength/8; i++) {
*mDest = _m_por(*mSrc1, *mSrc2);
mSrc1++;
mSrc2++;
mDest++;
}
_m_empty();
#endif
return (0);
#else
return (-1);
#endif
}


int SDL_imageFilterBitOr(unsigned char *Src1, unsigned char *Src2, unsigned char *Dest, unsigned int length)
{
unsigned int i, istart;
unsigned char *cursrc1, *cursrc2, *curdst;


if ((Src1 == NULL) || (Src2 == NULL) || (Dest == NULL))
return(-1);
if (length == 0)
return(0);

if ((SDL_imageFilterMMXdetect()) && (length > 7)) {


SDL_imageFilterBitOrMMX(Src1, Src2, Dest, length);


if ((length & 7) > 0) {

istart = length & 0xfffffff8;
cursrc1 = &Src1[istart];
cursrc2 = &Src2[istart];
curdst = &Dest[istart];
} else {

return (0);
}
} else {

istart = 0;
cursrc1 = Src1;
cursrc2 = Src2;
curdst = Dest;
}


for (i = istart; i < length; i++) {
*curdst = *cursrc1 | *cursrc2;

cursrc1++;
cursrc2++;
curdst++;
}
return (0);
}


static int SDL_imageFilterDivASM(unsigned char *Src1, unsigned char *Src2, unsigned char *Dest, unsigned int SrcLength)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
pusha
mov edx, Src1
mov esi, Src2
mov edi, Dest
mov ecx, SrcLength
align 16
L10191:
mov bl, [esi]
cmp bl, 0
jnz L10192
mov [edi], 255
jmp  L10193
L10192:
xor ah, ah
mov al, [edx]
div   bl
mov [edi], al
L10193:
inc edx
inc esi
inc edi
dec ecx
jnz L10191
popa
}
#else



asm volatile (
#  if defined(i386)
"pushl %%ebx \n\t"
".align 16     \n\t"
"1: mov (%%esi), %%bl  \n\t"
"cmp       $0, %%bl    \n\t"
"jnz 2f                \n\t"
"movb  $255, (%%edi)   \n\t"
"jmp 3f                \n\t"
"2: xor %%ah, %%ah     \n\t"
"mov   (%%edx), %%al   \n\t"
"div   %%bl            \n\t"
"mov   %%al, (%%edi)   \n\t"
"3: inc %%edx          \n\t"
"inc %%esi \n\t"
"inc %%edi \n\t"
"dec %%ecx \n\t"
"jnz 1b    \n\t"
"popl %%ebx \n\t"
: "+d" (Src1),
"+S" (Src2),
"+c" (SrcLength),
"+D" (Dest)
:
: "memory", "rax"
#  elif defined(__x86_64__)
".align 16     \n\t"
"1: mov (%%rsi), %%bl  \n\t"
"cmp       $0, %%bl    \n\t"
"jnz 2f                \n\t"
"movb  $255, (%%rdi)   \n\t"
"jmp 3f                \n\t"
"2: xor %%ah, %%ah     \n\t"
"mov   (%%rdx), %%al   \n\t"
"div   %%bl            \n\t"
"mov   %%al, (%%rdi)   \n\t"
"3: inc %%rdx          \n\t"
"inc %%rsi \n\t"
"inc %%rdi \n\t"
"dec %%rcx \n\t"
"jnz 1b    \n\t"
: "+d" (Src1),
"+S" (Src2),
"+c" (SrcLength),
"+D" (Dest)
:
: "memory", "rax", "rbx"
#  endif
);
#endif
return (0);
#else
return (-1);
#endif
}


int SDL_imageFilterDiv(unsigned char *Src1, unsigned char *Src2, unsigned char *Dest, unsigned int length)
{
unsigned int i, istart;
unsigned char *cursrc1, *cursrc2, *curdst;


if ((Src1 == NULL) || (Src2 == NULL) || (Dest == NULL))
return(-1);
if (length == 0)
return(0);

if (SDL_imageFilterMMXdetect()) {
if (length > 0) {

SDL_imageFilterDivASM(Src1, Src2, Dest, length);


return (0);
} else {
return (-1);
}
}


istart = 0;
cursrc1 = Src1;
cursrc2 = Src2;
curdst = Dest;














for (i = istart; i < length; i++) {
if (*cursrc2 == 0) {
*curdst = 255;
} else {
*curdst = (int)*cursrc1 / (int)*cursrc2;  // (int) for efficiency
}

cursrc1++;
cursrc2++;
curdst++;
}

return (0);
}




static int SDL_imageFilterBitNegationMMX(unsigned char *Src1, unsigned char *Dest, unsigned int SrcLength)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
pusha
pcmpeqb mm1, mm1
mov eax, Src1
mov edi, Dest
mov ecx, SrcLength
shr ecx,  3
align 16
L91117:
movq mm0, [eax]
pxor mm0, mm1
movq [edi], mm0
add eax, 8
add edi,  8
dec ecx
jnz L91117
emms
popa
}
#else

__m64 *mSrc1 = (__m64*)Src1;
__m64 *mDest = (__m64*)Dest;
__m64 mm1;
mm1 = _m_pcmpeqb(mm1, mm1);
int i;
for (i = 0; i < SrcLength/8; i++) {
*mDest = _m_pxor(*mSrc1, mm1);
mSrc1++;
mDest++;
}
_m_empty();

#endif
return (0);
#else
return (-1);
#endif
}


int SDL_imageFilterBitNegation(unsigned char *Src1, unsigned char *Dest, unsigned int length)
{
unsigned int i, istart;
unsigned char *cursrc1, *curdst;


if ((Src1 == NULL) || (Dest == NULL))
return(-1);
if (length == 0)
return(0);

if ((SDL_imageFilterMMXdetect()) && (length > 7)) {

SDL_imageFilterBitNegationMMX(Src1, Dest, length);


if ((length & 7) > 0) {

istart = length & 0xfffffff8;
cursrc1 = &Src1[istart];
curdst = &Dest[istart];
} else {

return (0);
}
} else {

istart = 0;
cursrc1 = Src1;
curdst = Dest;
}


for (i = istart; i < length; i++) {
*curdst = ~(*cursrc1);

cursrc1++;
curdst++;
}

return (0);
}


static int SDL_imageFilterAddByteMMX(unsigned char *Src1, unsigned char *Dest, unsigned int SrcLength, unsigned char C)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
pusha

mov al, C
mov ah, al
mov bx, ax
shl eax, 16
mov ax, bx
movd mm1, eax
movd mm2, eax
punpckldq mm1, mm2
mov eax, Src1
mov edi, Dest
mov ecx, SrcLength
shr ecx,  3
align 16
L1021:
movq mm0, [eax]
paddusb mm0,  mm1
movq [edi], mm0
add eax, 8
add edi, 8
dec              ecx
jnz             L1021
emms
popa
}
#else

__m64 *mSrc1 = (__m64*)Src1;
__m64 *mDest = (__m64*)Dest;

int i;
memset(&i, C, 4);
__m64 mm1 = _m_from_int(i);
__m64 mm2 = _m_from_int(i);
mm1 = _m_punpckldq(mm1, mm2);
//__m64 mm1 = _m_from_int64(lli); // x86_64 only
for (i = 0; i < SrcLength/8; i++) {
*mDest = _m_paddusb(*mSrc1, mm1);
mSrc1++;
mDest++;
}
_m_empty();
#endif
return (0);
#else
return (-1);
#endif
}


int SDL_imageFilterAddByte(unsigned char *Src1, unsigned char *Dest, unsigned int length, unsigned char C)
{
unsigned int i, istart;
int iC;
unsigned char *cursrc1, *curdest;
int result;


if ((Src1 == NULL) || (Dest == NULL))
return(-1);
if (length == 0)
return(0);


if (C == 0) {
memcpy(Src1, Dest, length);
return (0);
}

if ((SDL_imageFilterMMXdetect()) && (length > 7)) {


SDL_imageFilterAddByteMMX(Src1, Dest, length, C);


if ((length & 7) > 0) {

istart = length & 0xfffffff8;
cursrc1 = &Src1[istart];
curdest = &Dest[istart];
} else {

return (0);
}
} else {

istart = 0;
cursrc1 = Src1;
curdest = Dest;
}


iC = (int) C;
for (i = istart; i < length; i++) {
result = (int) *cursrc1 + iC;
if (result > 255)
result = 255;
*curdest = (unsigned char) result;

cursrc1++;
curdest++;
}
return (0);
}


static int SDL_imageFilterAddUintMMX(unsigned char *Src1, unsigned char *Dest, unsigned int SrcLength, unsigned int C, unsigned int D)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
pusha

mov eax, C
movd mm1, eax
mov eax, D
movd mm2, eax
punpckldq mm1, mm2
mov eax, Src1
mov edi, Dest
mov ecx, SrcLength
shr ecx,  3
align 16
L11023:
movq mm0, [eax]
paddusb mm0,  mm1
movq [edi],  mm0
add eax, 8
add edi, 8
dec              ecx
jnz             L11023
emms
popa
}
#else

__m64 *mSrc1 = (__m64*)Src1;
__m64 *mDest = (__m64*)Dest;

__m64 mm1 = _m_from_int(C);
__m64 mm2 = _m_from_int(C);
mm1 = _m_punpckldq(mm1, mm2);
//__m64 mm1 = _m_from_int64(lli); // x86_64 only
int i;
for (i = 0; i < SrcLength/8; i++) {
*mDest = _m_paddusb(*mSrc1, mm1);
mSrc1++;
mDest++;
}
_m_empty();
#endif
return (0);
#else
return (-1);
#endif
}


int SDL_imageFilterAddUint(unsigned char *Src1, unsigned char *Dest, unsigned int length, unsigned int C)
{
unsigned int i, j, istart, D;
int iC[4];
unsigned char *cursrc1;
unsigned char *curdest;
int result;


if ((Src1 == NULL) || (Dest == NULL))
return(-1);
if (length == 0)
return(0);


if (C == 0) {
memcpy(Src1, Dest, length);
return (0);
}

if ((SDL_imageFilterMMXdetect()) && (length > 7)) {


D=SWAP_32(C);
SDL_imageFilterAddUintMMX(Src1, Dest, length, C, D);


if ((length & 7) > 0) {

istart = length & 0xfffffff8;
cursrc1 = &Src1[istart];
curdest = &Dest[istart];
} else {

return (0);
}
} else {

istart = 0;
cursrc1 = Src1;
curdest = Dest;
}


iC[3] = (int) ((C >> 24) & 0xff);
iC[2] = (int) ((C >> 16) & 0xff);
iC[1] = (int) ((C >>  8) & 0xff);
iC[0] = (int) ((C >>  0) & 0xff);
for (i = istart; i < length; i += 4) {
for (j = 0; j < 4; j++) {
if ((i+j)<length) {
result = (int) *cursrc1 + iC[j];
if (result > 255) result = 255;
*curdest = (unsigned char) result;

cursrc1++;
curdest++;
}
}
}
return (0);
}


static int SDL_imageFilterAddByteToHalfMMX(unsigned char *Src1, unsigned char *Dest, unsigned int SrcLength, unsigned char C,
unsigned char *Mask)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
pusha

mov al, C
mov ah, al
mov bx, ax
shl eax, 16
mov ax, bx
movd mm1, eax
movd mm2, eax
punpckldq mm1, mm2
mov edx, Mask
movq mm0, [edx]
mov eax, Src1
mov edi, Dest
mov ecx,  SrcLength
shr ecx,  3
align 16
L1022:
movq mm2, [eax]
psrlw mm2, 1
pand mm2, mm0
paddusb mm2,  mm1
movq [edi], mm2
add eax, 8
add edi, 8
dec              ecx
jnz             L1022
emms
popa
}
#else

__m64 *mSrc1 = (__m64*)Src1;
__m64 *mDest = (__m64*)Dest;
__m64 *mMask = (__m64*)Mask;

int i;
memset(&i, C, 4);
__m64 mm1 = _m_from_int(i);
__m64 mm2 = _m_from_int(i);
mm1 = _m_punpckldq(mm1, mm2);
//__m64 mm1 = _m_from_int64(lli); // x86_64 only
for (i = 0; i < SrcLength/8; i++) {
__m64 mm2 = _m_psrlwi(*mSrc1, 1);
mm2 = _m_pand(mm2, *mMask);

*mDest = _m_paddusb(mm1, mm2);
mSrc1++;
mDest++;
}
_m_empty();
#endif
return (0);
#else
return (-1);
#endif
}


int SDL_imageFilterAddByteToHalf(unsigned char *Src1, unsigned char *Dest, unsigned int length, unsigned char C)
{
static unsigned char Mask[8] = { 0x7F, 0x7F, 0x7F, 0x7F, 0x7F, 0x7F, 0x7F, 0x7F };
unsigned int i, istart;
int iC;
unsigned char *cursrc1;
unsigned char *curdest;
int result;


if ((Src1 == NULL) || (Dest == NULL))
return(-1);
if (length == 0)
return(0);

if ((SDL_imageFilterMMXdetect()) && (length > 7)) {


SDL_imageFilterAddByteToHalfMMX(Src1, Dest, length, C, Mask);


if ((length & 7) > 0) {

istart = length & 0xfffffff8;
cursrc1 = &Src1[istart];
curdest = &Dest[istart];
} else {

return (0);
}
} else {

istart = 0;
cursrc1 = Src1;
curdest = Dest;
}


iC = (int) C;
for (i = istart; i < length; i++) {
result = (int) (*cursrc1 / 2) + iC;
if (result > 255)
result = 255;
*curdest = (unsigned char) result;

cursrc1++;
curdest++;
}

return (0);
}


int SDL_imageFilterSubByteMMX(unsigned char *Src1, unsigned char *Dest, unsigned int SrcLength, unsigned char C)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
pusha

mov al, C
mov ah, al
mov bx, ax
shl eax, 16
mov ax, bx
movd mm1, eax
movd mm2, eax
punpckldq mm1, mm2
mov eax, Src1
mov edi, Dest
mov ecx,  SrcLength
shr ecx,  3
align 16
L1023:
movq mm0, [eax]
psubusb mm0,  mm1
movq [edi], mm0
add eax, 8
add edi, 8
dec              ecx
jnz             L1023
emms
popa
}
#else

__m64 *mSrc1 = (__m64*)Src1;
__m64 *mDest = (__m64*)Dest;

int i;
memset(&i, C, 4);
__m64 mm1 = _m_from_int(i);
__m64 mm2 = _m_from_int(i);
mm1 = _m_punpckldq(mm1, mm2);
//__m64 mm1 = _m_from_int64(lli); // x86_64 only
for (i = 0; i < SrcLength/8; i++) {
*mDest = _m_psubusb(*mSrc1, mm1);
mSrc1++;
mDest++;
}
_m_empty();
#endif
return (0);
#else
return (-1);
#endif
}


int SDL_imageFilterSubByte(unsigned char *Src1, unsigned char *Dest, unsigned int length, unsigned char C)
{
unsigned int i, istart;
int iC;
unsigned char *cursrc1;
unsigned char *curdest;
int result;


if ((Src1 == NULL) || (Dest == NULL))
return(-1);
if (length == 0)
return(0);


if (C == 0) {
memcpy(Src1, Dest, length);
return (0);
}

if ((SDL_imageFilterMMXdetect()) && (length > 7)) {


SDL_imageFilterSubByteMMX(Src1, Dest, length, C);


if ((length & 7) > 0) {

istart = length & 0xfffffff8;
cursrc1 = &Src1[istart];
curdest = &Dest[istart];
} else {

return (0);
}
} else {

istart = 0;
cursrc1 = Src1;
curdest = Dest;
}


iC = (int) C;
for (i = istart; i < length; i++) {
result = (int) *cursrc1 - iC;
if (result < 0)
result = 0;
*curdest = (unsigned char) result;

cursrc1++;
curdest++;
}
return (0);
}


static int SDL_imageFilterSubUintMMX(unsigned char *Src1, unsigned char *Dest, unsigned int SrcLength, unsigned int C, unsigned int D)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
pusha

mov eax, C
movd mm1, eax
mov eax, D
movd mm2, eax
punpckldq mm1, mm2
mov eax, Src1
mov edi, Dest
mov ecx,  SrcLength
shr ecx,  3
align 16
L11024:
movq mm0, [eax]
psubusb mm0, mm1
movq [edi], mm0
add eax, 8
add edi, 8
dec              ecx
jnz             L11024
emms
popa
}
#else

__m64 *mSrc1 = (__m64*)Src1;
__m64 *mDest = (__m64*)Dest;

__m64 mm1 = _m_from_int(C);
__m64 mm2 = _m_from_int(C);
mm1 = _m_punpckldq(mm1, mm2);
//__m64 mm1 = _m_from_int64(lli); // x86_64 only
int i;
for (i = 0; i < SrcLength/8; i++) {
*mDest = _m_psubusb(*mSrc1, mm1);
mSrc1++;
mDest++;
}
_m_empty();
#endif
return (0);
#else
return (-1);
#endif
}


int SDL_imageFilterSubUint(unsigned char *Src1, unsigned char *Dest, unsigned int length, unsigned int C)
{
unsigned int i, j, istart, D;
int iC[4];
unsigned char *cursrc1;
unsigned char *curdest;
int result;


if ((Src1 == NULL) || (Dest == NULL))
return(-1);
if (length == 0)
return(0);


if (C == 0) {
memcpy(Src1, Dest, length);
return (0);
}

if ((SDL_imageFilterMMXdetect()) && (length > 7)) {


D=SWAP_32(C);
SDL_imageFilterSubUintMMX(Src1, Dest, length, C, D);


if ((length & 7) > 0) {

istart = length & 0xfffffff8;
cursrc1 = &Src1[istart];
curdest = &Dest[istart];
} else {

return (0);
}
} else {

istart = 0;
cursrc1 = Src1;
curdest = Dest;
}


iC[3] = (int) ((C >> 24) & 0xff);
iC[2] = (int) ((C >> 16) & 0xff);
iC[1] = (int) ((C >>  8) & 0xff);
iC[0] = (int) ((C >>  0) & 0xff);
for (i = istart; i < length; i += 4) {
for (j = 0; j < 4; j++) {
if ((i+j)<length) {
result = (int) *cursrc1 - iC[j];
if (result < 0) result = 0;
*curdest = (unsigned char) result;

cursrc1++;
curdest++;
}
}
}
return (0);
}


static int SDL_imageFilterShiftRightMMX(unsigned char *Src1, unsigned char *Dest, unsigned int SrcLength, unsigned char N,
unsigned char *Mask)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
pusha
mov edx, Mask
movq mm0, [edx]
xor ecx, ecx
mov cl,  N
movd mm3,  ecx
pcmpeqb mm1, mm1
L10240:
psrlw mm1,  1
pand mm1, mm0

dec               cl
jnz            L10240

mov eax, Src1
mov edi, Dest
mov ecx,  SrcLength
shr ecx,  3
align 16
L10241:
movq mm0, [eax]
psrlw mm0, mm3
pand mm0, mm1

movq [edi], mm0
add eax, 8
add edi, 8
dec              ecx
jnz            L10241
emms
popa
}
#else

__m64 *mSrc1 = (__m64*)Src1;
__m64 *mDest = (__m64*)Dest;
__m64 *mMask = (__m64*)Mask;
__m64 mm1;
int i;
mm1 = _m_pcmpeqb(mm1, mm1);

for (i = 0; i < N; i++) {
mm1 = _m_psrlwi(mm1, 1);
mm1 = _m_pand(mm1, *mMask);
}

for (i = 0; i < SrcLength/8; i++) {
__m64 mm0 = _m_psrlwi(*mSrc1, N);
*mDest = _m_pand(mm0, mm1);
mSrc1++;
mDest++;
}
_m_empty();
#endif
return (0);
#else
return (-1);
#endif
}


int SDL_imageFilterShiftRight(unsigned char *Src1, unsigned char *Dest, unsigned int length, unsigned char N)
{
static unsigned char Mask[8] = { 0x7F, 0x7F, 0x7F, 0x7F, 0x7F, 0x7F, 0x7F, 0x7F };
unsigned int i, istart;
unsigned char *cursrc1;
unsigned char *curdest;


if ((Src1 == NULL) || (Dest == NULL))
return(-1);
if (length == 0)
return(0);


if (N > 8) {
return (-1);
}


if (N == 0) {
memcpy(Src1, Dest, length);
return (0);
}

if ((SDL_imageFilterMMXdetect()) && (length > 7)) {


SDL_imageFilterShiftRightMMX(Src1, Dest, length, N, Mask);


if ((length & 7) > 0) {

istart = length & 0xfffffff8;
cursrc1 = &Src1[istart];
curdest = &Dest[istart];
} else {

return (0);
}
} else {

istart = 0;
cursrc1 = Src1;
curdest = Dest;
}


for (i = istart; i < length; i++) {
*curdest = (unsigned char) *cursrc1 >> N;

cursrc1++;
curdest++;
}

return (0);
}


static int SDL_imageFilterShiftRightUintMMX(unsigned char *Src1, unsigned char *Dest, unsigned int SrcLength, unsigned char N)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
pusha
mov eax, Src1
mov edi, Dest
mov ecx, SrcLength
shr ecx, 3
align 16
L13023:
movq mm0, [eax]
psrld mm0, N
movq [edi], mm0
add eax, 8
add edi, 8
dec              ecx
jnz             L13023
emms
popa
}
#else

__m64 *mSrc1 = (__m64*)Src1;
__m64 *mDest = (__m64*)Dest;
int i;
for (i = 0; i < SrcLength/8; i++) {
*mDest = _m_psrldi(*mSrc1, N);
mSrc1++;
mDest++;
}
_m_empty();
#endif
return (0);
#else
return (-1);
#endif
}


int SDL_imageFilterShiftRightUint(unsigned char *Src1, unsigned char *Dest, unsigned int length, unsigned char N)
{
unsigned int i, istart;
unsigned char *cursrc1, *curdest;
unsigned int *icursrc1, *icurdest;
unsigned int result;


if ((Src1 == NULL) || (Dest == NULL))
return(-1);
if (length == 0)
return(0);

if (N > 32) {
return (-1);
}


if (N == 0) {
memcpy(Src1, Dest, length);
return (0);
}

if ((SDL_imageFilterMMXdetect()) && (length > 7)) {

SDL_imageFilterShiftRightUintMMX(Src1, Dest, length, N);


if ((length & 7) > 0) {

istart = length & 0xfffffff8;
cursrc1 = &Src1[istart];
curdest = &Dest[istart];
} else {

return (0);
}
} else {

istart = 0;
cursrc1 = Src1;
curdest = Dest;
}


icursrc1=(unsigned int *)cursrc1;
icurdest=(unsigned int *)curdest;
for (i = istart; i < length; i += 4) {
if ((i+4)<length) {
result = ((unsigned int)*icursrc1 >> N);
*icurdest = result;
}

icursrc1++;
icurdest++;
}

return (0);
}


static int SDL_imageFilterMultByByteMMX(unsigned char *Src1, unsigned char *Dest, unsigned int SrcLength, unsigned char C)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
pusha

mov al, C
xor ah, ah
mov bx, ax
shl eax, 16
mov ax, bx
movd mm1, eax
movd mm2, eax
punpckldq mm1, mm2
pxor mm0, mm0
mov eax, Src1
mov edi, Dest
mov ecx, SrcLength
shr ecx, 3
cmp al, 128
jg             L10251
align 16
L10250:
movq mm3, [eax]
movq mm4, mm3
punpcklbw mm3, mm0
punpckhbw mm4, mm0
pmullw mm3, mm1
pmullw mm4, mm1
packuswb mm3, mm4
movq [edi], mm3
add eax, 8
add edi, 8
dec              ecx
jnz            L10250
jmp            L10252
align 16
L10251:
movq mm3, [eax]
movq mm4, mm3
punpcklbw mm3, mm0
punpckhbw mm4, mm0
pmullw mm3, mm1
pmullw mm4, mm1

movq mm5, mm3
movq mm6, mm4
psraw mm5, 15
psraw mm6, 15
pxor mm3, mm5
pxor mm4, mm6
psubsw mm3, mm5
psubsw mm4, mm6
packuswb mm3, mm4
movq [edi], mm3
add eax, 8
add edi, 8
dec              ecx
jnz            L10251
L10252:
emms
popa
}
#else

__m64 *mSrc1 = (__m64*)Src1;
__m64 *mDest = (__m64*)Dest;
__m64 mm0 = _m_from_int(0);

int i;
i = C | C<<16;
__m64 mm1 = _m_from_int(i);
__m64 mm2 = _m_from_int(i);
mm1 = _m_punpckldq(mm1, mm2);
// long long lli = C | C<<16 | (long long)C<<32 | (long long)C<<48;
//__m64 mm1 = _m_from_int64(lli); // x86_64 only
if (C <= 128) {
for (i = 0; i < SrcLength/8; i++) {
__m64 mm3, mm4;
mm3 = _m_punpcklbw(*mSrc1, mm0);
mm4 = _m_punpckhbw(*mSrc1, mm0);
mm3 = _m_pmullw(mm3, mm1);
mm4 = _m_pmullw(mm4, mm1);
*mDest = _m_packuswb(mm3, mm4);
mSrc1++;
mDest++;
}
} else {
for (i = 0; i < SrcLength/8; i++) {
__m64 mm3, mm4, mm5, mm6;
mm3 = _m_punpcklbw(*mSrc1, mm0);
mm4 = _m_punpckhbw(*mSrc1, mm0);
mm3 = _m_pmullw(mm3, mm1);
mm4 = _m_pmullw(mm4, mm1);

mm5 = _m_psrawi(mm3, 15);
mm6 = _m_psrawi(mm4, 15);
mm3 = _m_pxor(mm3, mm5);
mm4 = _m_pxor(mm4, mm6);
mm3 = _m_psubsw(mm3, mm5);
mm4 = _m_psubsw(mm4, mm6);
*mDest = _m_packuswb(mm3, mm4);
mSrc1++;
mDest++;
}
}
_m_empty();
#endif
return (0);
#else
return (-1);
#endif
}


int SDL_imageFilterMultByByte(unsigned char *Src1, unsigned char *Dest, unsigned int length, unsigned char C)
{
unsigned int i, istart;
int iC;
unsigned char *cursrc1;
unsigned char *curdest;
int result;


if ((Src1 == NULL) || (Dest == NULL))
return(-1);
if (length == 0)
return(0);


if (C == 1) {
memcpy(Src1, Dest, length);
return (0);
}

if ((SDL_imageFilterMMXdetect()) && (length > 7)) {

SDL_imageFilterMultByByteMMX(Src1, Dest, length, C);


if ((length & 7) > 0) {

istart = length & 0xfffffff8;
cursrc1 = &Src1[istart];
curdest = &Dest[istart];
} else {

return (0);
}
} else {

istart = 0;
cursrc1 = Src1;
curdest = Dest;
}


iC = (int) C;
for (i = istart; i < length; i++) {
result = (int) *cursrc1 * iC;
if (result > 255)
result = 255;
*curdest = (unsigned char) result;

cursrc1++;
curdest++;
}

return (0);
}


static int SDL_imageFilterShiftRightAndMultByByteMMX(unsigned char *Src1, unsigned char *Dest, unsigned int SrcLength, unsigned char N,
unsigned char C)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
pusha

mov al, C
xor ah, ah
mov bx, ax
shl eax, 16
mov ax, bx
movd mm1, eax
movd mm2, eax
punpckldq mm1, mm2
xor ecx, ecx
mov cl, N
movd mm7, ecx
pxor mm0, mm0
mov eax, Src1
mov edi, Dest
mov ecx, SrcLength
shr ecx, 3
align 16
L1026:
movq mm3, [eax]
movq mm4, mm3
punpcklbw mm3, mm0
punpckhbw mm4, mm0
psrlw mm3, mm7
psrlw mm4, mm7
pmullw mm3, mm1
pmullw mm4, mm1
packuswb mm3, mm4
movq [edi], mm3
add eax, 8
add edi, 8
dec              ecx
jnz             L1026
emms
popa
}
#else

__m64 *mSrc1 = (__m64*)Src1;
__m64 *mDest = (__m64*)Dest;
__m64 mm0 = _m_from_int(0);

int i;
i = (C<<16)|C;
__m64 mm1 = _m_from_int(i);
__m64 mm2 = _m_from_int(i);
mm1 = _m_punpckldq(mm1, mm2);
for (i = 0; i < SrcLength/8; i++) {
__m64 mm3, mm4, mm5, mm6;
mm3 = _m_punpcklbw(*mSrc1, mm0);
mm4 = _m_punpckhbw(*mSrc1, mm0);
mm3 = _m_psrlwi(mm3, N);
mm4 = _m_psrlwi(mm4, N);
mm3 = _m_pmullw(mm3, mm1);
mm4 = _m_pmullw(mm4, mm1);
*mDest = _m_packuswb(mm3, mm4);
mSrc1++;
mDest++;
}
_m_empty();
#endif
return (0);
#else
return (-1);
#endif
}


int SDL_imageFilterShiftRightAndMultByByte(unsigned char *Src1, unsigned char *Dest, unsigned int length, unsigned char N,
unsigned char C)
{
unsigned int i, istart;
int iC;
unsigned char *cursrc1;
unsigned char *curdest;
int result;


if ((Src1 == NULL) || (Dest == NULL))
return(-1);
if (length == 0)
return(0);


if (N > 8) {
return (-1);
}


if ((N == 0) && (C == 1)) {
memcpy(Src1, Dest, length);
return (0);
}

if ((SDL_imageFilterMMXdetect()) && (length > 7)) {

SDL_imageFilterShiftRightAndMultByByteMMX(Src1, Dest, length, N, C);


if ((length & 7) > 0) {

istart = length & 0xfffffff8;
cursrc1 = &Src1[istart];
curdest = &Dest[istart];
} else {

return (0);
}
} else {

istart = 0;
cursrc1 = Src1;
curdest = Dest;
}


iC = (int) C;
for (i = istart; i < length; i++) {
result = (int) (*cursrc1 >> N) * iC;
if (result > 255)
result = 255;
*curdest = (unsigned char) result;

cursrc1++;
curdest++;
}

return (0);
}


static int SDL_imageFilterShiftLeftByteMMX(unsigned char *Src1, unsigned char *Dest, unsigned int SrcLength, unsigned char N,
unsigned char *Mask)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
pusha
mov edx, Mask
movq mm0, [edx]
xor ecx, ecx
mov cl, N
movd mm3, ecx
pcmpeqb mm1, mm1
L10270:
psllw mm1, 1
pand mm1, mm0

dec cl
jnz            L10270

mov eax, Src1
mov edi, Dest
mov ecx, SrcLength
shr ecx, 3
align 16
L10271:
movq mm0, [eax]
psllw mm0, mm3
pand mm0, mm1

movq [edi], mm0
add eax, 8
add edi, 8
dec              ecx
jnz            L10271
emms
popa
}
#else

__m64 *mSrc1 = (__m64*)Src1;
__m64 *mDest = (__m64*)Dest;
__m64 *mMask = (__m64*)Mask;
__m64 mm1;
int i;
mm1 = _m_pcmpeqb(mm1, mm1);

for (i = 0; i < N; i++) {
mm1 = _m_psllwi(mm1, 1);
mm1 = _m_pand(mm1, *mMask);
}

for (i = 0; i < SrcLength/8; i++) {
__m64 mm0 = _m_psllwi(*mSrc1, N);
*mDest = _m_pand(mm0, mm1);
mSrc1++;
mDest++;
}
_m_empty();
#endif
return (0);
#else
return (-1);
#endif
}


int SDL_imageFilterShiftLeftByte(unsigned char *Src1, unsigned char *Dest, unsigned int length, unsigned char N)
{
static unsigned char Mask[8] = { 0xFE, 0xFE, 0xFE, 0xFE, 0xFE, 0xFE, 0xFE, 0xFE };
unsigned int i, istart;
unsigned char *cursrc1, *curdest;
int result;


if ((Src1 == NULL) || (Dest == NULL))
return(-1);
if (length == 0)
return(0);

if (N > 8) {
return (-1);
}


if (N == 0) {
memcpy(Src1, Dest, length);
return (0);
}

if ((SDL_imageFilterMMXdetect()) && (length > 7)) {

SDL_imageFilterShiftLeftByteMMX(Src1, Dest, length, N, Mask);


if ((length & 7) > 0) {

istart = length & 0xfffffff8;
cursrc1 = &Src1[istart];
curdest = &Dest[istart];
} else {

return (0);
}
} else {

istart = 0;
cursrc1 = Src1;
curdest = Dest;
}


for (i = istart; i < length; i++) {
result = ((int) *cursrc1 << N) & 0xff;
*curdest = (unsigned char) result;

cursrc1++;
curdest++;
}

return (0);
}


static int SDL_imageFilterShiftLeftUintMMX(unsigned char *Src1, unsigned char *Dest, unsigned int SrcLength, unsigned char N)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
pusha
mov eax, Src1
mov edi, Dest
mov ecx, SrcLength
shr ecx, 3
align 16
L12023:
movq mm0, [eax]
pslld mm0, N
movq [edi], mm0
add eax, 8
add edi, 8
dec              ecx
jnz             L12023
emms
popa
}
#else

__m64 *mSrc1 = (__m64*)Src1;
__m64 *mDest = (__m64*)Dest;
int i;
for (i = 0; i < SrcLength/8; i++) {
*mDest = _m_pslldi(*mSrc1, N);
mSrc1++;
mDest++;
}
_m_empty();
#endif
return (0);
#else
return (-1);
#endif
}


int SDL_imageFilterShiftLeftUint(unsigned char *Src1, unsigned char *Dest, unsigned int length, unsigned char N)
{
unsigned int i, istart;
unsigned char *cursrc1, *curdest;
unsigned int *icursrc1, *icurdest;
unsigned int result;


if ((Src1 == NULL) || (Dest == NULL))
return(-1);
if (length == 0)
return(0);

if (N > 32) {
return (-1);
}


if (N == 0) {
memcpy(Src1, Dest, length);
return (0);
}

if ((SDL_imageFilterMMXdetect()) && (length > 7)) {

SDL_imageFilterShiftLeftUintMMX(Src1, Dest, length, N);


if ((length & 7) > 0) {

istart = length & 0xfffffff8;
cursrc1 = &Src1[istart];
curdest = &Dest[istart];
} else {

return (0);
}
} else {

istart = 0;
cursrc1 = Src1;
curdest = Dest;
}


icursrc1=(unsigned int *)cursrc1;
icurdest=(unsigned int *)curdest;
for (i = istart; i < length; i += 4) {
if ((i+4)<length) {
result = ((unsigned int)*icursrc1 << N);
*icurdest = result;
}

icursrc1++;
icurdest++;
}

return (0);
}


static int SDL_imageFilterShiftLeftMMX(unsigned char *Src1, unsigned char *Dest, unsigned int SrcLength, unsigned char N)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
pusha
xor eax, eax
mov al, N
movd mm7, eax
pxor mm0, mm0
mov eax, Src1
mov edi, Dest
mov ecx, SrcLength
shr ecx, 3
cmp al, 7
jg             L10281
align 16
L10280:
movq mm3, [eax]
movq mm4, mm3
punpcklbw mm3, mm0
punpckhbw mm4, mm0
psllw mm3, mm7
psllw mm4, mm7
packuswb mm3, mm4
movq [edi], mm3
add eax, 8
add edi, 8
dec              ecx
jnz            L10280
jmp            L10282
align 16
L10281:
movq mm3, [eax]
movq mm4, mm3
punpcklbw mm3, mm0
punpckhbw mm4, mm0
psllw mm3, mm7
psllw mm4, mm7

movq mm5, mm3
movq mm6, mm4
psraw mm5, 15
psraw mm6, 15
pxor mm3, mm5
pxor mm4, mm6
psubsw mm3, mm5
psubsw mm4, mm6
packuswb mm3, mm4
movq [edi], mm3
add eax, 8
add edi, 8
dec              ecx
jnz            L10281
L10282:
emms
popa
}
#else

__m64 *mSrc1 = (__m64*)Src1;
__m64 *mDest = (__m64*)Dest;
__m64 mm0 = _m_from_int(0);
int i;
if (N <= 7) {
for (i = 0; i < SrcLength/8; i++) {
__m64 mm3, mm4;
mm3 = _m_punpcklbw(*mSrc1, mm0);
mm4 = _m_punpckhbw(*mSrc1, mm0);
mm3 = _m_psllwi(mm3, N);
mm4 = _m_psllwi(mm4, N);
*mDest = _m_packuswb(mm3, mm4);
mSrc1++;
mDest++;
}
} else {
for (i = 0; i < SrcLength/8; i++) {
__m64 mm3, mm4, mm5, mm6;
mm3 = _m_punpcklbw(*mSrc1, mm0);
mm4 = _m_punpckhbw(*mSrc1, mm0);
mm3 = _m_psllwi(mm3, N);
mm4 = _m_psllwi(mm4, N);

mm5 = _m_psrawi(mm3, 15);
mm6 = _m_psrawi(mm4, 15);
mm3 = _m_pxor(mm3, mm5);
mm4 = _m_pxor(mm4, mm6);
mm3 = _m_psubsw(mm3, mm5);
mm4 = _m_psubsw(mm4, mm6);
*mDest = _m_packuswb(mm3, mm4);
mSrc1++;
mDest++;
}
}
_m_empty();
#endif
return (0);
#else
return (-1);
#endif
}


int SDL_imageFilterShiftLeft(unsigned char *Src1, unsigned char *Dest, unsigned int length, unsigned char N)
{
unsigned int i, istart;
unsigned char *cursrc1, *curdest;
int result;


if ((Src1 == NULL) || (Dest == NULL))
return(-1);
if (length == 0)
return(0);

if (N > 8) {
return (-1);
}


if (N == 0) {
memcpy(Src1, Dest, length);
return (0);
}

if ((SDL_imageFilterMMXdetect()) && (length > 7)) {

SDL_imageFilterShiftLeftMMX(Src1, Dest, length, N);


if ((length & 7) > 0) {

istart = length & 0xfffffff8;
cursrc1 = &Src1[istart];
curdest = &Dest[istart];
} else {

return (0);
}
} else {

istart = 0;
cursrc1 = Src1;
curdest = Dest;
}


for (i = istart; i < length; i++) {
result = (int) *cursrc1 << N;
if (result > 255)
result = 255;
*curdest = (unsigned char) result;

cursrc1++;
curdest++;
}

return (0);
}


static int SDL_imageFilterBinarizeUsingThresholdMMX(unsigned char *Src1, unsigned char *Dest, unsigned int SrcLength, unsigned char T)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
pusha

pcmpeqb mm1, mm1
pcmpeqb mm2, mm2
mov al, T
mov ah, al
mov bx, ax
shl eax, 16
mov ax, bx
movd mm3, eax
movd mm4, eax
punpckldq mm3, mm4
psubusb mm2, mm3
mov eax, Src1
mov edi, Dest
mov ecx, SrcLength
shr ecx, 3
align 16
L1029:
movq mm0, [eax]
paddusb mm0, mm2
pcmpeqb mm0, mm1
movq [edi], mm0
add eax, 8
add edi, 8
dec              ecx
jnz             L1029
emms
popa
}
#else

__m64 *mSrc1 = (__m64*)Src1;
__m64 *mDest = (__m64*)Dest;

__m64 mm1 = _m_pcmpeqb(mm1, mm1);
__m64 mm2 = _m_pcmpeqb(mm2, mm2);
int i;
memset(&i, T, 4);
__m64 mm3 = _m_from_int(i);
__m64 mm4 = _m_from_int(i);
mm3 = _m_punpckldq(mm3, mm4);
mm2 = _m_psubusb(mm2, mm3);
//__m64 mm3 = _m_from_int64(lli); // x86_64 only
for (i = 0; i < SrcLength/8; i++) {
__m64 mm0 = _m_paddusb(*mSrc1, mm2);
*mDest = _m_pcmpeqb(mm0, mm1);
mSrc1++;
mDest++;
}
_m_empty();
#endif
return (0);
#else
return (-1);
#endif
}


int SDL_imageFilterBinarizeUsingThreshold(unsigned char *Src1, unsigned char *Dest, unsigned int length, unsigned char T)
{
unsigned int i, istart;
unsigned char *cursrc1;
unsigned char *curdest;


if ((Src1 == NULL) || (Dest == NULL))
return(-1);
if (length == 0)
return(0);


if (T == 0) {
memset(Dest, 255, length);
return (0);
}

if ((SDL_imageFilterMMXdetect()) && (length > 7)) {

SDL_imageFilterBinarizeUsingThresholdMMX(Src1, Dest, length, T);


if ((length & 7) > 0) {

istart = length & 0xfffffff8;
cursrc1 = &Src1[istart];
curdest = &Dest[istart];
} else {

return (0);
}
} else {

istart = 0;
cursrc1 = Src1;
curdest = Dest;
}


for (i = istart; i < length; i++) {
*curdest = (unsigned char)(((unsigned char)*cursrc1 >= T) ? 255 : 0);

cursrc1++;
curdest++;
}

return (0);
}


static int SDL_imageFilterClipToRangeMMX(unsigned char *Src1, unsigned char *Dest, unsigned int SrcLength, unsigned char Tmin,
unsigned char Tmax)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
pusha
pcmpeqb mm1, mm1

mov al, Tmax
mov ah, al
mov bx, ax
shl eax, 16
mov ax, bx
movd mm3, eax
movd mm4, eax
punpckldq mm3, mm4
psubusb mm1, mm3

mov al, Tmin
mov ah, al
mov bx, ax
shl eax, 16
mov ax, bx
movd mm5, eax
movd mm4, eax
punpckldq mm5, mm4
movq mm7, mm5
paddusb mm7, mm1
mov eax, Src1
mov edi, Dest
mov ecx, SrcLength
shr ecx, 3
align 16
L1030:
movq mm0, [eax]
paddusb mm0, mm1
psubusb mm0, mm7
paddusb mm0, mm5
movq [edi], mm0
add eax, 8
add edi, 8
dec              ecx
jnz             L1030
emms
popa
}
#else

__m64 *mSrc1 = (__m64*)Src1;
__m64 *mDest = (__m64*)Dest;
__m64 mm1 = _m_pcmpeqb(mm1, mm1);
int i;

__m64 mm3, mm4;
memset(&i, Tmax, 4);
mm3 = _m_from_int(i);
mm4 = _m_from_int(i);
mm3 = _m_punpckldq(mm3, mm4);
mm1 = _m_psubusb(mm1, mm3);
//__m64 mm3 = _m_from_int64(lli); // x86_64 only

__m64 mm5, mm7;
memset(&i, Tmin, 4);
mm5 = _m_from_int(i);
mm4 = _m_from_int(i);
mm5 = _m_punpckldq(mm5, mm4);
mm7 = _m_paddusb(mm5, mm1);
for (i = 0; i < SrcLength/8; i++) {
__m64 mm0;
mm0 = _m_paddusb(*mSrc1, mm1);
mm0 = _m_psubusb(mm0, mm7);
*mDest = _m_paddusb(mm0, mm5);
mSrc1++;
mDest++;
}
_m_empty();
#endif
return (0);
#else
return (-1);
#endif
}


int SDL_imageFilterClipToRange(unsigned char *Src1, unsigned char *Dest, unsigned int length, unsigned char Tmin,
unsigned char Tmax)
{
unsigned int i, istart;
unsigned char *cursrc1;
unsigned char *curdest;


if ((Src1 == NULL) || (Dest == NULL))
return(-1);
if (length == 0)
return(0);


if ((Tmin == 0) && (Tmax == 25)) {
memcpy(Src1, Dest, length);
return (0);
}

if ((SDL_imageFilterMMXdetect()) && (length > 7)) {

SDL_imageFilterClipToRangeMMX(Src1, Dest, length, Tmin, Tmax);


if ((length & 7) > 0) {

istart = length & 0xfffffff8;
cursrc1 = &Src1[istart];
curdest = &Dest[istart];
} else {

return (0);
}
} else {

istart = 0;
cursrc1 = Src1;
curdest = Dest;
}


for (i = istart; i < length; i++) {
if (*cursrc1 < Tmin) {
*curdest = Tmin;
} else if (*cursrc1 > Tmax) {
*curdest = Tmax;
} else {
*curdest = *cursrc1;
}

cursrc1++;
curdest++;
}

return (0);
}


static int SDL_imageFilterNormalizeLinearMMX(unsigned char *Src1, unsigned char *Dest, unsigned int SrcLength, int Cmin, int Cmax,
int Nmin, int Nmax)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
pusha
mov ax, WORD PTR Nmax
mov bx, WORD PTR Cmax
sub ax, WORD PTR Nmin
sub bx, WORD PTR Cmin
jz             L10311
xor dx, dx
div               bx
jmp            L10312
L10311:
mov ax, 255
L10312:
mov bx, ax
shl eax, 16
mov ax, bx
movd mm0, eax
movd mm1, eax
punpckldq mm0, mm1

mov ax, WORD PTR Cmin
mov bx, ax
shl eax, 16
mov ax, bx
movd mm1, eax
movd mm2, eax
punpckldq mm1, mm2

mov ax, WORD PTR Nmin
mov bx, ax
shl eax, 16
mov ax, bx
movd mm2, eax
movd mm3, eax
punpckldq mm2, mm3
pxor mm7, mm7
mov eax, Src1
mov edi, Dest
mov ecx, SrcLength
shr ecx, 3
align 16
L1031:
movq mm3, [eax]
movq mm4, mm3
punpcklbw mm3, mm7
punpckhbw mm4, mm7
psubusb mm3, mm1
psubusb mm4, mm1
pmullw mm3, mm0
pmullw mm4, mm0
paddusb mm3, mm2
paddusb mm4, mm2

movq mm5, mm3
movq mm6, mm4
psraw mm5, 15
psraw mm6, 15
pxor mm3, mm5
pxor mm4, mm6
psubsw mm3, mm5
psubsw mm4, mm6
packuswb mm3, mm4
movq [edi], mm3
add eax, 8
add edi, 8
dec              ecx
jnz             L1031
emms
popa
}
#else

__m64 *mSrc1 = (__m64*)Src1;
__m64 *mDest = (__m64*)Dest;
__m64 mm0, mm1, mm2, mm3;

int i;

unsigned short a = Nmax - Nmin;
unsigned short b = Cmax - Cmin;
if (b == 0) {
a = 255;
} else {
a /= b;
}
i = (a<<16)|a;
mm0 = _m_from_int(i);
mm1 = _m_from_int(i);
mm0 = _m_punpckldq(mm0, mm1);

i = (Cmin<<16)|(short)Cmin;
mm1 = _m_from_int(i);
mm2 = _m_from_int(i);
mm1 = _m_punpckldq(mm1, mm2);

i = (Nmin<<16)|(short)Nmin;
mm2 = _m_from_int(i);
mm3 = _m_from_int(i);
mm2 = _m_punpckldq(mm2, mm3);
__m64 mm7 = _m_from_int(0);
for (i = 0; i < SrcLength/8; i++) {
__m64 mm3, mm4, mm5, mm6;
mm3 = _m_punpcklbw(*mSrc1, mm7);
mm4 = _m_punpckhbw(*mSrc1, mm7);
mm3 = _m_psubusb(mm3, mm1);
mm4 = _m_psubusb(mm4, mm1);
mm3 = _m_pmullw(mm3, mm0);
mm4 = _m_pmullw(mm4, mm0);
mm3 = _m_paddusb(mm3, mm2);
mm4 = _m_paddusb(mm4, mm2);

mm5 = _m_psrawi(mm3, 15);
mm6 = _m_psrawi(mm4, 15);
mm3 = _m_pxor(mm3, mm5);
mm4 = _m_pxor(mm4, mm6);
mm3 = _m_psubsw(mm3, mm5);
mm4 = _m_psubsw(mm4, mm6);
*mDest = _m_packuswb(mm3, mm4);
mSrc1++;
mDest++;
}
_m_empty();
#endif
return (0);
#else
return (-1);
#endif
}


int SDL_imageFilterNormalizeLinear(unsigned char *Src, unsigned char *Dest, unsigned int length, int Cmin, int Cmax, int Nmin,
int Nmax)
{
unsigned int i, istart;
unsigned char *cursrc;
unsigned char *curdest;
int dN, dC, factor;
int result;


if ((Src == NULL) || (Dest == NULL))
return(-1);
if (length == 0)
return(0);

if ((SDL_imageFilterMMXdetect()) && (length > 7)) {

SDL_imageFilterNormalizeLinearMMX(Src, Dest, length, Cmin, Cmax, Nmin, Nmax);


if ((length & 7) > 0) {

istart = length & 0xfffffff8;
cursrc = &Src[istart];
curdest = &Dest[istart];
} else {

return (0);
}
} else {

istart = 0;
cursrc = Src;
curdest = Dest;
}


dC = Cmax - Cmin;
if (dC == 0)
return (0);
dN = Nmax - Nmin;
factor = dN / dC;
for (i = istart; i < length; i++) {
result = factor * ((int) (*cursrc) - Cmin) + Nmin;
if (result > 255)
result = 255;
*curdest = (unsigned char) result;

cursrc++;
curdest++;
}

return (0);
}




int SDL_imageFilterConvolveKernel3x3Divide(unsigned char *Src, unsigned char *Dest, int rows, int columns,
signed short *Kernel, unsigned char Divisor)
{

if ((Src == NULL) || (Dest == NULL) || (Kernel == NULL))
return(-1);

if ((columns < 3) || (rows < 3) || (Divisor == 0))
return (-1);

if ((SDL_imageFilterMMXdetect())) {
//#ifdef USE_MMX
#if defined(USE_MMX) && defined(i386)
#if !defined(GCC__)
__asm
{
pusha
pxor mm0, mm0
xor ebx, ebx
mov bl, Divisor
mov edx, Kernel
movq mm5, [edx]
add edx, 8
movq mm6, [edx]
add edx, 8
movq mm7, [edx]

mov eax, columns
mov esi, Src
mov edi, Dest
add edi, eax
inc              edi
mov edx, rows
sub edx, 2

L10320:
mov ecx, eax
sub ecx, 2
align 16
L10322:

movq mm1, [esi]
add esi, eax
movq mm2, [esi]
add esi, eax
movq mm3, [esi]
punpcklbw mm1, mm0
punpcklbw mm2, mm0
punpcklbw mm3, mm0
pmullw mm1, mm5
pmullw mm2, mm6
pmullw mm3, mm7
paddsw mm1, mm2
paddsw mm1, mm3
movq mm2, mm1
psrlq mm1, 32
paddsw mm1, mm2
movq mm3, mm1
psrlq mm1, 16
paddsw mm1, mm3

movd mm2, eax
movd mm3, edx
movd eax, mm1
psraw mm1, 15
movd edx, mm1
idiv bx
movd mm1, eax
packuswb mm1, mm0
movd eax, mm1
mov [edi], al
movd edx, mm3
movd eax, mm2

sub esi, eax
sub esi, eax
inc              esi
inc              edi

dec              ecx
jnz            L10322
add esi, 2
add edi, 2
dec              edx
jnz            L10320

emms
popa
}
#else
asm volatile
("pusha		     \n\t" "pxor      %%mm0, %%mm0 \n\t"
"xor       %%ebx, %%ebx \n\t"
"mov           %5, %%bl \n\t"
"mov          %4, %%edx \n\t"
"movq    (%%edx), %%mm5 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm6 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm7 \n\t"

"mov          %3, %%eax \n\t"
"mov          %1, %%esi \n\t"
"mov          %0, %%edi \n\t"
"add       %%eax, %%edi \n\t"
"inc              %%edi \n\t"
"mov          %2, %%edx \n\t"
"sub          $2, %%edx \n\t"

".L10320:               \n\t" "mov       %%eax, %%ecx \n\t"
"sub          $2, %%ecx \n\t"
".align 16              \n\t"
".L10322:               \n\t"

"movq    (%%esi), %%mm1 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%esi), %%mm2 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%esi), %%mm3 \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpcklbw %%mm0, %%mm2 \n\t"
"punpcklbw %%mm0, %%mm3 \n\t"
"pmullw    %%mm5, %%mm1 \n\t"
"pmullw    %%mm6, %%mm2 \n\t"
"pmullw    %%mm7, %%mm3 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm3, %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"psrlq       $32, %%mm1 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"movq      %%mm1, %%mm3 \n\t"
"psrlq       $16, %%mm1 \n\t"
"paddsw    %%mm3, %%mm1 \n\t"

"movd      %%eax, %%mm2 \n\t"
"movd      %%edx, %%mm3 \n\t"
"movd      %%mm1, %%eax \n\t"
"psraw       $15, %%mm1 \n\t"
"movd      %%mm1, %%edx \n\t"
"idivw             %%bx \n\t"
"movd      %%eax, %%mm1 \n\t"
"packuswb  %%mm0, %%mm1 \n\t"
"movd      %%mm1, %%eax \n\t"
"mov      %%al, (%%edi) \n\t"
"movd      %%mm3, %%edx \n\t"
"movd      %%mm2, %%eax \n\t"

"sub       %%eax, %%esi \n\t"
"sub       %%eax, %%esi \n\t"
"inc              %%esi \n\t"
"inc              %%edi \n\t"

"dec              %%ecx \n\t"
"jnz            .L10322 \n\t"
"add          $2, %%esi \n\t"
"add          $2, %%edi \n\t"
"dec              %%edx \n\t"
"jnz            .L10320 \n\t"

"emms                   \n\t"
"popa                   \n\t":"=m" (Dest)
:"m"(Src),
"m"(rows),
"m"(columns),
"m"(Kernel),
"m"(Divisor)
);
#endif
#endif
return (0);
} else {

return (-1);
}
}


int SDL_imageFilterConvolveKernel5x5Divide(unsigned char *Src, unsigned char *Dest, int rows, int columns,
signed short *Kernel, unsigned char Divisor)
{

if ((Src == NULL) || (Dest == NULL) || (Kernel == NULL))
return(-1);

if ((columns < 5) || (rows < 5) || (Divisor == 0))
return (-1);

if ((SDL_imageFilterMMXdetect())) {
//#ifdef USE_MMX
#if defined(USE_MMX) && defined(i386)
#if !defined(GCC__)
__asm
{
pusha
pxor mm0, mm0
xor ebx, ebx
mov bl, Divisor
movd mm5, ebx
mov edx, Kernel
mov esi, Src
mov edi, Dest
add edi, 2
mov eax, columns
shl eax, 1
add edi, eax
shr eax, 1
mov ebx, rows
sub ebx, 4

L10330:
mov ecx, eax
sub ecx, 4
align 16
L10332:
pxor mm7, mm7
movd mm6, esi

movq mm1, [esi]
movq mm2, mm1
add esi, eax
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
add esi, eax
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
add esi, eax
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
add esi, eax
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
punpcklbw mm1, mm0
punpckhbw mm2, mm0
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1

movq mm3, mm7
psrlq mm7, 32
paddsw mm7, mm3
movq mm2, mm7
psrlq mm7, 16
paddsw mm7, mm2

movd mm1, eax
movd mm2, ebx
movd mm3, edx
movd eax, mm7
psraw mm7, 15
movd ebx, mm5
movd edx, mm7
idiv bx
movd mm7, eax
packuswb mm7, mm0
movd eax, mm7
mov [edi], al
movd edx, mm3
movd ebx, mm2
movd eax, mm1

movd esi, mm6
sub edx, 72
inc              esi
inc              edi

dec              ecx
jnz            L10332
add esi, 4
add edi, 4
dec              ebx
jnz            L10330

emms
popa
}
#else
asm volatile
("pusha		     \n\t" "pxor      %%mm0, %%mm0 \n\t"
"xor       %%ebx, %%ebx \n\t"
"mov           %5, %%bl \n\t"
"movd      %%ebx, %%mm5 \n\t"
"mov          %4, %%edx \n\t"
"mov          %1, %%esi \n\t"
"mov          %0, %%edi \n\t"
"add          $2, %%edi \n\t"
"mov          %3, %%eax \n\t"
"shl          $1, %%eax \n\t"
"add       %%eax, %%edi \n\t"
"shr          $1, %%eax \n\t"
"mov          %2, %%ebx \n\t"
"sub          $4, %%ebx \n\t"

".L10330:               \n\t" "mov       %%eax, %%ecx \n\t"
"sub          $4, %%ecx \n\t"
".align 16              \n\t"
".L10332:               \n\t" "pxor      %%mm7, %%mm7 \n\t"
"movd      %%esi, %%mm6 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq      %%mm7, %%mm3 \n\t"
"psrlq       $32, %%mm7 \n\t"
"paddsw    %%mm3, %%mm7 \n\t"
"movq      %%mm7, %%mm2 \n\t"
"psrlq       $16, %%mm7 \n\t"
"paddsw    %%mm2, %%mm7 \n\t"

"movd      %%eax, %%mm1 \n\t"
"movd      %%ebx, %%mm2 \n\t"
"movd      %%edx, %%mm3 \n\t"
"movd      %%mm7, %%eax \n\t"
"psraw       $15, %%mm7 \n\t"
"movd      %%mm5, %%ebx \n\t"
"movd      %%mm7, %%edx \n\t"
"idivw             %%bx \n\t"
"movd      %%eax, %%mm7 \n\t"
"packuswb  %%mm0, %%mm7 \n\t"
"movd      %%mm7, %%eax \n\t"
"mov      %%al, (%%edi) \n\t"
"movd      %%mm3, %%edx \n\t"
"movd      %%mm2, %%ebx \n\t"
"movd      %%mm1, %%eax \n\t"

"movd      %%mm6, %%esi \n\t"
"sub         $72, %%edx \n\t"
"inc              %%esi \n\t"
"inc              %%edi \n\t"

"dec              %%ecx \n\t"
"jnz            .L10332 \n\t"
"add          $4, %%esi \n\t"
"add          $4, %%edi \n\t"
"dec              %%ebx \n\t"
"jnz            .L10330 \n\t"

"emms                   \n\t"
"popa                   \n\t":"=m" (Dest)
:"m"(Src),
"m"(rows),
"m"(columns),
"m"(Kernel),
"m"(Divisor)
);
#endif
#endif
return (0);
} else {

return (-1);
}
}


int SDL_imageFilterConvolveKernel7x7Divide(unsigned char *Src, unsigned char *Dest, int rows, int columns,
signed short *Kernel, unsigned char Divisor)
{

if ((Src == NULL) || (Dest == NULL) || (Kernel == NULL))
return(-1);

if ((columns < 7) || (rows < 7) || (Divisor == 0))
return (-1);

if ((SDL_imageFilterMMXdetect())) {
//#ifdef USE_MMX
#if defined(USE_MMX) && defined(i386)
#if !defined(GCC__)
__asm
{
pusha
pxor mm0, mm0
xor ebx, ebx
mov bl, Divisor
movd mm5, ebx
mov edx, Kernel
mov esi, Src
mov edi, Dest
add edi, 3
mov eax, columns
add edi, eax
add edi, eax
add edi, eax
mov ebx, rows
sub ebx, 6

L10340:
mov ecx, eax
sub ecx, 6
align 16
L10342:
pxor mm7, mm7
movd mm6, esi

movq mm1, [esi]
movq mm2, mm1
add esi, eax
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
add esi, eax
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
add esi, eax
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
add esi, eax
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
add esi, eax
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
add esi, eax
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
punpcklbw mm1, mm0
punpckhbw mm2, mm0
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1

movq mm3, mm7
psrlq mm7, 32
paddsw mm7, mm3
movq mm2, mm7
psrlq mm7, 16
paddsw mm7, mm2

movd mm1, eax
movd mm2, ebx
movd mm3, edx
movd eax, mm7
psraw mm7, 15
movd ebx, mm5
movd edx, mm7
idiv bx
movd mm7, eax
packuswb mm7, mm0
movd eax, mm7
mov [edi], al
movd edx, mm3
movd ebx, mm2
movd eax, mm1

movd esi, mm6
sub edx, 104
inc              esi
inc              edi

dec              ecx
jnz            L10342
add esi, 6
add edi, 6
dec              ebx
jnz            L10340

emms
popa
}
#else
asm volatile
("pusha		     \n\t" "pxor      %%mm0, %%mm0 \n\t"
"xor       %%ebx, %%ebx \n\t"
"mov           %5, %%bl \n\t"
"movd      %%ebx, %%mm5 \n\t"
"mov          %4, %%edx \n\t"
"mov          %1, %%esi \n\t"
"mov          %0, %%edi \n\t"
"add          $3, %%edi \n\t"
"mov          %3, %%eax \n\t"
"add       %%eax, %%edi \n\t"
"add       %%eax, %%edi \n\t" "add       %%eax, %%edi \n\t" "mov          %2, %%ebx \n\t"
"sub          $6, %%ebx \n\t"

".L10340:               \n\t" "mov       %%eax, %%ecx \n\t"
"sub          $6, %%ecx \n\t"
".align 16              \n\t"
".L10342:               \n\t" "pxor      %%mm7, %%mm7 \n\t"
"movd      %%esi, %%mm6 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq      %%mm7, %%mm3 \n\t"
"psrlq       $32, %%mm7 \n\t"
"paddsw    %%mm3, %%mm7 \n\t"
"movq      %%mm7, %%mm2 \n\t"
"psrlq       $16, %%mm7 \n\t"
"paddsw    %%mm2, %%mm7 \n\t"

"movd      %%eax, %%mm1 \n\t"
"movd      %%ebx, %%mm2 \n\t"
"movd      %%edx, %%mm3 \n\t"
"movd      %%mm7, %%eax \n\t"
"psraw       $15, %%mm7 \n\t"
"movd      %%mm5, %%ebx \n\t"
"movd      %%mm7, %%edx \n\t"
"idivw             %%bx \n\t"
"movd      %%eax, %%mm7 \n\t"
"packuswb  %%mm0, %%mm7 \n\t"
"movd      %%mm7, %%eax \n\t"
"mov      %%al, (%%edi) \n\t"
"movd      %%mm3, %%edx \n\t"
"movd      %%mm2, %%ebx \n\t"
"movd      %%mm1, %%eax \n\t"

"movd      %%mm6, %%esi \n\t"
"sub        $104, %%edx \n\t"
"inc              %%esi \n\t"
"inc              %%edi \n\t"

"dec              %%ecx \n\t"
"jnz            .L10342 \n\t"
"add          $6, %%esi \n\t"
"add          $6, %%edi \n\t"
"dec              %%ebx \n\t"
"jnz            .L10340 \n\t"

"emms                   \n\t"
"popa                   \n\t":"=m" (Dest)
:"m"(Src),
"m"(rows),
"m"(columns),
"m"(Kernel),
"m"(Divisor)
);
#endif
#endif
return (0);
} else {

return (-1);
}
}


int SDL_imageFilterConvolveKernel9x9Divide(unsigned char *Src, unsigned char *Dest, int rows, int columns,
signed short *Kernel, unsigned char Divisor)
{

if ((Src == NULL) || (Dest == NULL) || (Kernel == NULL))
return(-1);

if ((columns < 9) || (rows < 9) || (Divisor == 0))
return (-1);

if ((SDL_imageFilterMMXdetect())) {
//#ifdef USE_MMX
#if defined(USE_MMX) && defined(i386)
#if !defined(GCC__)
__asm
{
pusha
pxor mm0, mm0
xor ebx, ebx
mov bl, Divisor
movd mm5, ebx
mov edx, Kernel
mov esi, Src
mov edi, Dest
add edi, 4
mov eax, columns
add edi, eax
add edi, eax
add edi, eax
add edi, eax
mov ebx, rows
sub ebx, 8

L10350:
mov ecx, eax
sub ecx, 8
align 16
L10352:
pxor mm7, mm7
movd mm6, esi

movq mm1, [esi]
movq mm2, mm1
inc              esi
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1
movq mm1, [esi]
dec              esi
add esi, eax
movq mm3, [edx]
add edx, 8
punpcklbw mm1, mm0
pmullw mm1, mm3
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
inc              esi
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1
movq mm1, [esi]
dec              esi
add esi, eax
movq mm3, [edx]
add edx, 8
punpcklbw mm1, mm0
pmullw mm1, mm3
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
inc              esi
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1
movq mm1, [esi]
dec              esi
add esi, eax
movq mm3, [edx]
add edx, 8
punpcklbw mm1, mm0
pmullw mm1, mm3
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
inc              esi
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1
movq mm1, [esi]
dec              esi
add esi, eax
movq mm3, [edx]
add edx, 8
punpcklbw mm1, mm0
pmullw mm1, mm3
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
inc              esi
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1
movq mm1, [esi]
dec              esi
add esi, eax
movq mm3, [edx]
add edx, 8
punpcklbw mm1, mm0
pmullw mm1, mm3
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
inc              esi
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1
movq mm1, [esi]
dec              esi
add esi, eax
movq mm3, [edx]
add edx, 8
punpcklbw mm1, mm0
pmullw mm1, mm3
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
inc              esi
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1
movq mm1, [esi]
dec              esi
add esi, eax
movq mm3, [edx]
add edx, 8
punpcklbw mm1, mm0
pmullw mm1, mm3
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
inc              esi
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1
movq mm1, [esi]
dec              esi
add esi, eax
movq mm3, [edx]
add edx, 8
punpcklbw mm1, mm0
pmullw mm1, mm3
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
inc              esi
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1
movq mm1, [esi]
movq mm3, [edx]
punpcklbw mm1, mm0
pmullw mm1, mm3
paddsw mm7, mm1

movq mm3, mm7
psrlq mm7, 32
paddsw mm7, mm3
movq mm2, mm7
psrlq mm7, 16
paddsw mm7, mm2

movd mm1, eax
movd mm2, ebx
movd mm3, edx
movd eax, mm7
psraw mm7, 15
movd ebx, mm5
movd edx, mm7
idiv bx
movd mm7, eax
packuswb mm7, mm0
movd eax, mm7
mov [edi], al
movd edx, mm3
movd ebx, mm2
movd eax, mm1

movd esi, mm6
sub edx, 208
inc              esi
inc              edi

dec              ecx
jnz            L10352
add esi, 8
add edi, 8
dec              ebx
jnz            L10350

emms
popa
}
#else
asm volatile
("pusha		     \n\t" "pxor      %%mm0, %%mm0 \n\t"
"xor       %%ebx, %%ebx \n\t"
"mov           %5, %%bl \n\t"
"movd      %%ebx, %%mm5 \n\t"
"mov          %4, %%edx \n\t"
"mov          %1, %%esi \n\t"
"mov          %0, %%edi \n\t"
"add          $4, %%edi \n\t"
"mov          %3, %%eax \n\t"
"add       %%eax, %%edi \n\t"
"add       %%eax, %%edi \n\t" "add       %%eax, %%edi \n\t" "add       %%eax, %%edi \n\t" "mov          %2, %%ebx \n\t"
"sub          $8, %%ebx \n\t"

".L10350:               \n\t" "mov       %%eax, %%ecx \n\t"
"sub          $8, %%ecx \n\t"
".align 16              \n\t"
".L10352:               \n\t" "pxor      %%mm7, %%mm7 \n\t"
"movd      %%esi, %%mm6 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"inc              %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"
"movq    (%%esi), %%mm1 \n\t"
"dec              %%esi \n\t" "add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"inc              %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"
"movq    (%%esi), %%mm1 \n\t"
"dec              %%esi \n\t" "add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"inc              %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"
"movq    (%%esi), %%mm1 \n\t"
"dec              %%esi \n\t" "add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"inc              %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"
"movq    (%%esi), %%mm1 \n\t"
"dec              %%esi \n\t" "add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"inc              %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"
"movq    (%%esi), %%mm1 \n\t"
"dec              %%esi \n\t" "add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"inc              %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"
"movq    (%%esi), %%mm1 \n\t"
"dec              %%esi \n\t" "add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"inc              %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"
"movq    (%%esi), %%mm1 \n\t"
"dec              %%esi \n\t" "add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"inc              %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"
"movq    (%%esi), %%mm1 \n\t"
"dec              %%esi \n\t" "add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"inc              %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"
"movq    (%%esi), %%mm1 \n\t"
"movq    (%%edx), %%mm3 \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq      %%mm7, %%mm3 \n\t"
"psrlq       $32, %%mm7 \n\t"
"paddsw    %%mm3, %%mm7 \n\t"
"movq      %%mm7, %%mm2 \n\t"
"psrlq       $16, %%mm7 \n\t"
"paddsw    %%mm2, %%mm7 \n\t"

"movd      %%eax, %%mm1 \n\t"
"movd      %%ebx, %%mm2 \n\t"
"movd      %%edx, %%mm3 \n\t"
"movd      %%mm7, %%eax \n\t"
"psraw       $15, %%mm7 \n\t"
"movd      %%mm5, %%ebx \n\t"
"movd      %%mm7, %%edx \n\t"
"idivw             %%bx \n\t"
"movd      %%eax, %%mm7 \n\t"
"packuswb  %%mm0, %%mm7 \n\t"
"movd      %%mm7, %%eax \n\t"
"mov      %%al, (%%edi) \n\t"
"movd      %%mm3, %%edx \n\t"
"movd      %%mm2, %%ebx \n\t"
"movd      %%mm1, %%eax \n\t"

"movd      %%mm6, %%esi \n\t"
"sub        $208, %%edx \n\t"
"inc              %%esi \n\t"
"inc              %%edi \n\t"

"dec              %%ecx \n\t"
"jnz            .L10352 \n\t"
"add          $8, %%esi \n\t"
"add          $8, %%edi \n\t"
"dec              %%ebx \n\t"
"jnz            .L10350 \n\t"

"emms                   \n\t"
"popa                   \n\t":"=m" (Dest)
:"m"(Src),
"m"(rows),
"m"(columns),
"m"(Kernel),
"m"(Divisor)
);
#endif
#endif
return (0);
} else {

return (-1);
}
}


int SDL_imageFilterConvolveKernel3x3ShiftRight(unsigned char *Src, unsigned char *Dest, int rows, int columns,
signed short *Kernel, unsigned char NRightShift)
{

if ((Src == NULL) || (Dest == NULL) || (Kernel == NULL))
return(-1);

if ((columns < 3) || (rows < 3) || (NRightShift > 7))
return (-1);

if ((SDL_imageFilterMMXdetect())) {
//#ifdef USE_MMX
#if defined(USE_MMX) && defined(i386)
#if !defined(GCC__)
__asm
{
pusha
pxor mm0, mm0
xor ebx, ebx
mov bl, NRightShift
movd mm4, ebx
mov edx, Kernel
movq mm5, [edx]
add edx, 8
movq mm6, [edx]
add edx, 8
movq mm7, [edx]

mov eax, columns
mov esi, Src
mov edi, Dest
add edi, eax
inc              edi
mov edx, rows
sub edx, 2

L10360:
mov ecx, eax
sub ecx, 2
align 16
L10362:

movq mm1, [esi]
add esi, eax
movq mm2, [esi]
add esi, eax
movq mm3, [esi]
punpcklbw mm1, mm0
punpcklbw mm2, mm0
punpcklbw mm3, mm0
psrlw mm1, mm4
psrlw mm2, mm4
psrlw mm3, mm4
pmullw mm1, mm5
pmullw mm2, mm6
pmullw mm3, mm7
paddsw mm1, mm2
paddsw mm1, mm3
movq mm2, mm1
psrlq mm1, 32
paddsw mm1, mm2
movq mm3, mm1
psrlq mm1, 16
paddsw mm1, mm3
packuswb mm1, mm0
movd ebx, mm1
mov [edi], bl

sub esi, eax
sub esi, eax
inc              esi
inc              edi

dec              ecx
jnz            L10362
add esi, 2
add edi, 2
dec              edx
jnz            L10360

emms
popa
}
#else
asm volatile
("pusha		     \n\t" "pxor      %%mm0, %%mm0 \n\t"
"xor       %%ebx, %%ebx \n\t"
"mov           %5, %%bl \n\t"
"movd      %%ebx, %%mm4 \n\t"
"mov          %4, %%edx \n\t"
"movq    (%%edx), %%mm5 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm6 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm7 \n\t"

"mov          %3, %%eax \n\t"
"mov          %1, %%esi \n\t"
"mov          %0, %%edi \n\t"
"add       %%eax, %%edi \n\t"
"inc              %%edi \n\t"
"mov          %2, %%edx \n\t"
"sub          $2, %%edx \n\t"

".L10360:               \n\t" "mov       %%eax, %%ecx \n\t"
"sub          $2, %%ecx \n\t"
".align 16              \n\t"
".L10362:               \n\t"

"movq    (%%esi), %%mm1 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%esi), %%mm2 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%esi), %%mm3 \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpcklbw %%mm0, %%mm2 \n\t"
"punpcklbw %%mm0, %%mm3 \n\t"
"psrlw     %%mm4, %%mm1 \n\t"
"psrlw     %%mm4, %%mm2 \n\t"
"psrlw     %%mm4, %%mm3 \n\t"
"pmullw    %%mm5, %%mm1 \n\t"
"pmullw    %%mm6, %%mm2 \n\t"
"pmullw    %%mm7, %%mm3 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm3, %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"psrlq       $32, %%mm1 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"movq      %%mm1, %%mm3 \n\t"
"psrlq       $16, %%mm1 \n\t"
"paddsw    %%mm3, %%mm1 \n\t"
"packuswb  %%mm0, %%mm1 \n\t"
"movd      %%mm1, %%ebx \n\t"
"mov      %%bl, (%%edi) \n\t"

"sub       %%eax, %%esi \n\t"
"sub       %%eax, %%esi \n\t" "inc              %%esi \n\t"
"inc              %%edi \n\t"

"dec              %%ecx \n\t"
"jnz            .L10362 \n\t"
"add          $2, %%esi \n\t"
"add          $2, %%edi \n\t"
"dec              %%edx \n\t"
"jnz            .L10360 \n\t"

"emms                   \n\t"
"popa                   \n\t":"=m" (Dest)
:"m"(Src),
"m"(rows),
"m"(columns),
"m"(Kernel),
"m"(NRightShift)
);
#endif
#endif
return (0);
} else {

return (-1);
}
}


int SDL_imageFilterConvolveKernel5x5ShiftRight(unsigned char *Src, unsigned char *Dest, int rows, int columns,
signed short *Kernel, unsigned char NRightShift)
{

if ((Src == NULL) || (Dest == NULL) || (Kernel == NULL))
return(-1);

if ((columns < 5) || (rows < 5) || (NRightShift > 7))
return (-1);

if ((SDL_imageFilterMMXdetect())) {
//#ifdef USE_MMX
#if defined(USE_MMX) && defined(i386)
#if !defined(GCC__)
__asm
{
pusha
pxor mm0, mm0
xor ebx, ebx
mov bl, NRightShift
movd mm5, ebx
mov edx, Kernel
mov esi, Src
mov edi, Dest
add edi, 2
mov eax, columns
shl eax, 1
add edi, eax
shr eax, 1
mov ebx, rows
sub ebx, 4

L10370:
mov ecx, eax
sub ecx, 4
align 16
L10372:
pxor mm7, mm7
movd mm6, esi

movq mm1, [esi]
movq mm2, mm1
add esi, eax
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
psrlw mm1, mm5
psrlw mm2, mm5
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
add esi, eax
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
psrlw mm1, mm5
psrlw mm2, mm5
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
add esi, eax
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
psrlw mm1, mm5
psrlw mm2, mm5
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
add esi, eax
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
psrlw mm1, mm5
psrlw mm2, mm5
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
punpcklbw mm1, mm0
punpckhbw mm2, mm0
psrlw mm1, mm5
psrlw mm2, mm5
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1

movq mm3, mm7
psrlq mm7, 32
paddsw mm7, mm3
movq mm2, mm7
psrlq mm7, 16
paddsw mm7, mm2
movd mm1, eax
packuswb mm7, mm0
movd eax, mm7
mov [edi], al
movd eax, mm1

movd esi, mm6
sub edx, 72
inc              esi
inc              edi

dec              ecx
jnz            L10372
add esi, 4
add edi, 4
dec              ebx
jnz            L10370

emms
popa
}
#else
asm volatile
("pusha		     \n\t" "pxor      %%mm0, %%mm0 \n\t"
"xor       %%ebx, %%ebx \n\t"
"mov           %5, %%bl \n\t"
"movd      %%ebx, %%mm5 \n\t"
"mov          %4, %%edx \n\t"
"mov          %1, %%esi \n\t"
"mov          %0, %%edi \n\t"
"add          $2, %%edi \n\t"
"mov          %3, %%eax \n\t"
"shl          $1, %%eax \n\t"
"add       %%eax, %%edi \n\t"
"shr          $1, %%eax \n\t"
"mov          %2, %%ebx \n\t"
"sub          $4, %%ebx \n\t"

".L10370:               \n\t" "mov       %%eax, %%ecx \n\t"
"sub          $4, %%ecx \n\t"
".align 16              \n\t"
".L10372:               \n\t" "pxor      %%mm7, %%mm7 \n\t"
"movd      %%esi, %%mm6 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"psrlw     %%mm5, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"psrlw     %%mm5, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"psrlw     %%mm5, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"psrlw     %%mm5, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"psrlw     %%mm5, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq      %%mm7, %%mm3 \n\t"
"psrlq       $32, %%mm7 \n\t"
"paddsw    %%mm3, %%mm7 \n\t"
"movq      %%mm7, %%mm2 \n\t"
"psrlq       $16, %%mm7 \n\t"
"paddsw    %%mm2, %%mm7 \n\t"
"movd      %%eax, %%mm1 \n\t"
"packuswb  %%mm0, %%mm7 \n\t"
"movd      %%mm7, %%eax \n\t"
"mov      %%al, (%%edi) \n\t"
"movd      %%mm1, %%eax \n\t"

"movd      %%mm6, %%esi \n\t"
"sub         $72, %%edx \n\t"
"inc              %%esi \n\t"
"inc              %%edi \n\t"

"dec              %%ecx \n\t"
"jnz            .L10372 \n\t"
"add          $4, %%esi \n\t"
"add          $4, %%edi \n\t"
"dec              %%ebx \n\t"
"jnz            .L10370 \n\t"

"emms                   \n\t"
"popa                   \n\t":"=m" (Dest)
:"m"(Src),
"m"(rows),
"m"(columns),
"m"(Kernel),
"m"(NRightShift)
);
#endif
#endif
return (0);
} else {

return (-1);
}
}


int SDL_imageFilterConvolveKernel7x7ShiftRight(unsigned char *Src, unsigned char *Dest, int rows, int columns,
signed short *Kernel, unsigned char NRightShift)
{

if ((Src == NULL) || (Dest == NULL) || (Kernel == NULL))
return(-1);

if ((columns < 7) || (rows < 7) || (NRightShift > 7))
return (-1);

if ((SDL_imageFilterMMXdetect())) {
//#ifdef USE_MMX
#if defined(USE_MMX) && defined(i386)
#if !defined(GCC__)
__asm
{
pusha
pxor mm0, mm0
xor ebx, ebx
mov bl, NRightShift
movd mm5, ebx
mov edx, Kernel
mov esi, Src
mov edi, Dest
add edi, 3
mov eax, columns
add edi, eax
add edi, eax
add edi, eax
mov ebx, rows
sub ebx, 6

L10380:
mov ecx, eax
sub ecx, 6
align 16
L10382:
pxor mm7, mm7
movd mm6, esi

movq mm1, [esi]
movq mm2, mm1
add esi, eax
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
psrlw mm1, mm5
psrlw mm2, mm5
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
add esi, eax
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
psrlw mm1, mm5
psrlw mm2, mm5
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
add esi, eax
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
psrlw mm1, mm5
psrlw mm2, mm5
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
add esi, eax
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
psrlw mm1, mm5
psrlw mm2, mm5
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
add esi, eax
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
psrlw mm1, mm5
psrlw mm2, mm5
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
add esi, eax
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
psrlw mm1, mm5
psrlw mm2, mm5
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
punpcklbw mm1, mm0
punpckhbw mm2, mm0
psrlw mm1, mm5
psrlw mm2, mm5
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1

movq mm3, mm7
psrlq mm7, 32
paddsw mm7, mm3
movq mm2, mm7
psrlq mm7, 16
paddsw mm7, mm2
movd mm1, eax
packuswb mm7, mm0
movd eax, mm7
mov [edi], al
movd eax, mm1

movd esi, mm6
sub edx, 104
inc              esi
inc              edi

dec              ecx
jnz            L10382
add esi, 6
add edi, 6
dec              ebx
jnz            L10380

emms
popa
}
#else
asm volatile
("pusha		     \n\t" "pxor      %%mm0, %%mm0 \n\t"
"xor       %%ebx, %%ebx \n\t"
"mov           %5, %%bl \n\t"
"movd      %%ebx, %%mm5 \n\t"
"mov          %4, %%edx \n\t"
"mov          %1, %%esi \n\t"
"mov          %0, %%edi \n\t"
"add          $3, %%edi \n\t"
"mov          %3, %%eax \n\t"
"add       %%eax, %%edi \n\t"
"add       %%eax, %%edi \n\t" "add       %%eax, %%edi \n\t" "mov          %2, %%ebx \n\t"
"sub          $6, %%ebx \n\t"

".L10380:               \n\t" "mov       %%eax, %%ecx \n\t"
"sub          $6, %%ecx \n\t"
".align 16              \n\t"
".L10382:               \n\t" "pxor      %%mm7, %%mm7 \n\t"
"movd      %%esi, %%mm6 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"psrlw     %%mm5, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"psrlw     %%mm5, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"psrlw     %%mm5, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"psrlw     %%mm5, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"psrlw     %%mm5, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"psrlw     %%mm5, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"psrlw     %%mm5, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq      %%mm7, %%mm3 \n\t"
"psrlq       $32, %%mm7 \n\t"
"paddsw    %%mm3, %%mm7 \n\t"
"movq      %%mm7, %%mm2 \n\t"
"psrlq       $16, %%mm7 \n\t"
"paddsw    %%mm2, %%mm7 \n\t"
"movd      %%eax, %%mm1 \n\t"
"packuswb  %%mm0, %%mm7 \n\t"
"movd      %%mm7, %%eax \n\t"
"mov      %%al, (%%edi) \n\t"
"movd      %%mm1, %%eax \n\t"

"movd      %%mm6, %%esi \n\t"
"sub        $104, %%edx \n\t"
"inc              %%esi \n\t"
"inc              %%edi \n\t"

"dec              %%ecx \n\t"
"jnz            .L10382 \n\t"
"add          $6, %%esi \n\t"
"add          $6, %%edi \n\t"
"dec              %%ebx \n\t"
"jnz            .L10380 \n\t"

"emms                   \n\t"
"popa                   \n\t":"=m" (Dest)
:"m"(Src),
"m"(rows),
"m"(columns),
"m"(Kernel),
"m"(NRightShift)
);
#endif
#endif
return (0);
} else {

return (-1);
}
}


int SDL_imageFilterConvolveKernel9x9ShiftRight(unsigned char *Src, unsigned char *Dest, int rows, int columns,
signed short *Kernel, unsigned char NRightShift)
{

if ((Src == NULL) || (Dest == NULL) || (Kernel == NULL))
return(-1);

if ((columns < 9) || (rows < 9) || (NRightShift > 7))
return (-1);

if ((SDL_imageFilterMMXdetect())) {
//#ifdef USE_MMX
#if defined(USE_MMX) && defined(i386)
#if !defined(GCC__)
__asm
{
pusha
pxor mm0, mm0
xor ebx, ebx
mov bl, NRightShift
movd mm5, ebx
mov edx, Kernel
mov esi, Src
mov edi, Dest
add edi, 4
mov eax, columns
add edi, eax
add edi, eax
add edi, eax
add edi, eax
mov ebx, rows
sub ebx, 8

L10390:
mov ecx, eax
sub ecx, 8
align 16
L10392:
pxor mm7, mm7
movd mm6, esi

movq mm1, [esi]
movq mm2, mm1
inc              esi
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
psrlw mm1, mm5
psrlw mm2, mm5
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1
movq mm1, [esi]
dec              esi
add esi, eax
movq mm3, [edx]
add edx, 8
punpcklbw mm1, mm0
psrlw mm1, mm5
pmullw mm1, mm3
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
inc              esi
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
psrlw mm1, mm5
psrlw mm2, mm5
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1
movq mm1, [esi]
dec              esi
add esi, eax
movq mm3, [edx]
add edx, 8
punpcklbw mm1, mm0
psrlw mm1, mm5
pmullw mm1, mm3
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
inc              esi
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
psrlw mm1, mm5
psrlw mm2, mm5
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1
movq mm1, [esi]
dec              esi
add esi, eax
movq mm3, [edx]
add edx, 8
punpcklbw mm1, mm0
psrlw mm1, mm5
pmullw mm1, mm3
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
inc              esi
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
psrlw mm1, mm5
psrlw mm2, mm5
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1
movq mm1, [esi]
dec              esi
add esi, eax
movq mm3, [edx]
add edx, 8
punpcklbw mm1, mm0
psrlw mm1, mm5
pmullw mm1, mm3
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
inc              esi
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
psrlw mm1, mm5
psrlw mm2, mm5
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1
movq mm1, [esi]
dec              esi
add esi, eax
movq mm3, [edx]
add edx, 8
punpcklbw mm1, mm0
psrlw mm1, mm5
pmullw mm1, mm3
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
inc              esi
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
psrlw mm1, mm5
psrlw mm2, mm5
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1
movq mm1, [esi]
dec              esi
add esi, eax
movq mm3, [edx]
add edx, 8
punpcklbw mm1, mm0
psrlw mm1, mm5
pmullw mm1, mm3
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
inc              esi
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
psrlw mm1, mm5
psrlw mm2, mm5
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1
movq mm1, [esi]
dec              esi
add esi, eax
movq mm3, [edx]
add edx, 8
punpcklbw mm1, mm0
psrlw mm1, mm5
pmullw mm1, mm3
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
inc              esi
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
psrlw mm1, mm5
psrlw mm2, mm5
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1
movq mm1, [esi]
dec              esi
add esi, eax
movq mm3, [edx]
add edx, 8
punpcklbw mm1, mm0
psrlw mm1, mm5
pmullw mm1, mm3
paddsw mm7, mm1

movq mm1, [esi]
movq mm2, mm1
inc              esi
movq mm3, [edx]
add edx, 8
movq mm4, [edx]
add edx, 8
punpcklbw mm1, mm0
punpckhbw mm2, mm0
psrlw mm1, mm5
psrlw mm2, mm5
pmullw mm1, mm3
pmullw mm2, mm4
paddsw mm1, mm2
paddsw mm7, mm1
movq mm1, [esi]
movq mm3, [edx]
punpcklbw mm1, mm0
psrlw mm1, mm5
pmullw mm1, mm3
paddsw mm7, mm1

movq mm3, mm7
psrlq mm7, 32
paddsw mm7, mm3
movq mm2, mm7
psrlq mm7, 16
paddsw mm7, mm2
movd mm1, eax
packuswb mm7, mm0
movd eax, mm7
mov [edi], al
movd eax, mm1

movd esi, mm6
sub edx, 208
inc              esi
inc              edi

dec              ecx
jnz            L10392
add esi, 8
add edi, 8
dec              ebx
jnz            L10390

emms
popa
}
#else
asm volatile
("pusha		     \n\t" "pxor      %%mm0, %%mm0 \n\t"
"xor       %%ebx, %%ebx \n\t"
"mov           %5, %%bl \n\t"
"movd      %%ebx, %%mm5 \n\t"
"mov          %4, %%edx \n\t"
"mov          %1, %%esi \n\t"
"mov          %0, %%edi \n\t"
"add          $4, %%edi \n\t"
"mov          %3, %%eax \n\t"
"add       %%eax, %%edi \n\t"
"add       %%eax, %%edi \n\t" "add       %%eax, %%edi \n\t" "add       %%eax, %%edi \n\t" "mov          %2, %%ebx \n\t"
"sub          $8, %%ebx \n\t"

".L10390:               \n\t" "mov       %%eax, %%ecx \n\t"
"sub          $8, %%ecx \n\t"
".align 16              \n\t"
".L10392:               \n\t" "pxor      %%mm7, %%mm7 \n\t"
"movd      %%esi, %%mm6 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"inc              %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"psrlw     %%mm5, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"
"movq    (%%esi), %%mm1 \n\t"
"dec              %%esi \n\t" "add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"inc              %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"psrlw     %%mm5, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"
"movq    (%%esi), %%mm1 \n\t"
"dec              %%esi \n\t" "add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"inc              %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"psrlw     %%mm5, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"
"movq    (%%esi), %%mm1 \n\t"
"dec              %%esi \n\t" "add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"inc              %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"psrlw     %%mm5, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"
"movq    (%%esi), %%mm1 \n\t"
"dec              %%esi \n\t" "add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"inc              %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"psrlw     %%mm5, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"
"movq    (%%esi), %%mm1 \n\t"
"dec              %%esi \n\t" "add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"inc              %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"psrlw     %%mm5, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"
"movq    (%%esi), %%mm1 \n\t"
"dec              %%esi \n\t" "add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"inc              %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"psrlw     %%mm5, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"
"movq    (%%esi), %%mm1 \n\t"
"dec              %%esi \n\t" "add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"inc              %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"psrlw     %%mm5, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"
"movq    (%%esi), %%mm1 \n\t"
"dec              %%esi \n\t" "add       %%eax, %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq    (%%esi), %%mm1 \n\t"
"movq      %%mm1, %%mm2 \n\t"
"inc              %%esi \n\t"
"movq    (%%edx), %%mm3 \n\t"
"add          $8, %%edx \n\t"
"movq    (%%edx), %%mm4 \n\t"
"add          $8, %%edx \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"punpckhbw %%mm0, %%mm2 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"psrlw     %%mm5, %%mm2 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"pmullw    %%mm4, %%mm2 \n\t"
"paddsw    %%mm2, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"
"movq    (%%esi), %%mm1 \n\t"
"movq    (%%edx), %%mm3 \n\t"
"punpcklbw %%mm0, %%mm1 \n\t"
"psrlw     %%mm5, %%mm1 \n\t"
"pmullw    %%mm3, %%mm1 \n\t"
"paddsw    %%mm1, %%mm7 \n\t"

"movq      %%mm7, %%mm3 \n\t"
"psrlq       $32, %%mm7 \n\t"
"paddsw    %%mm3, %%mm7 \n\t"
"movq      %%mm7, %%mm2 \n\t"
"psrlq       $16, %%mm7 \n\t"
"paddsw    %%mm2, %%mm7 \n\t"
"movd      %%eax, %%mm1 \n\t"
"packuswb  %%mm0, %%mm7 \n\t"
"movd      %%mm7, %%eax \n\t"
"mov      %%al, (%%edi) \n\t"
"movd      %%mm1, %%eax \n\t"

"movd      %%mm6, %%esi \n\t"
"sub        $208, %%edx \n\t"
"inc              %%esi \n\t"
"inc              %%edi \n\t"

"dec              %%ecx \n\t"
"jnz            .L10392 \n\t"
"add          $8, %%esi \n\t"
"add          $8, %%edi \n\t"
"dec              %%ebx \n\t"
"jnz            .L10390 \n\t"

"emms                   \n\t"
"popa                   \n\t":"=m" (Dest)
:"m"(Src),
"m"(rows),
"m"(columns),
"m"(Kernel),
"m"(NRightShift)
);
#endif
#endif
return (0);
} else {

return (-1);
}
}




int SDL_imageFilterSobelX(unsigned char *Src, unsigned char *Dest, int rows, int columns)
{

if ((Src == NULL) || (Dest == NULL))
return(-1);

if ((columns < 8) || (rows < 3))
return (-1);

if ((SDL_imageFilterMMXdetect())) {
//#ifdef USE_MMX
#if defined(USE_MMX) && defined(i386)
#if !defined(GCC__)
__asm
{
pusha
pxor mm0, mm0
mov eax, columns

mov esi, Src
mov edi, Dest
add edi, eax
inc              edi
mov edx, rows
sub edx, 2

L10400:
mov ecx, eax
shr ecx, 3
mov ebx, esi
movd mm1, edi
align 16
L10402:

movq mm4, [esi]
movq mm5, mm4
add esi, 2
punpcklbw mm4, mm0
punpckhbw mm5, mm0
movq mm6, [esi]
movq mm7, mm6
sub esi, 2
punpcklbw mm6, mm0
punpckhbw mm7, mm0
add esi, eax
movq mm2, [esi]
movq mm3, mm2
add esi, 2
punpcklbw mm2, mm0
punpckhbw mm3, mm0
paddw mm4, mm2
paddw mm5, mm3
paddw mm4, mm2
paddw mm5, mm3
movq mm2, [esi]
movq mm3, mm2
sub esi, 2
punpcklbw mm2, mm0
punpckhbw mm3, mm0
paddw mm6, mm2
paddw mm7, mm3
paddw mm6, mm2
paddw mm7, mm3
add esi, eax
movq mm2, [esi]
movq mm3, mm2
add esi, 2
punpcklbw mm2, mm0
punpckhbw mm3, mm0
paddw mm4, mm2
paddw mm5, mm3
movq mm2, [esi]
movq mm3, mm2
sub esi, 2
punpcklbw mm2, mm0
punpckhbw mm3, mm0
paddw mm6, mm2
paddw mm7, mm3

movq mm2, mm4
psrlq mm4, 32
psubw mm4, mm2
movq mm3, mm6
psrlq mm6, 32
psubw mm6, mm3
punpckldq mm4, mm6
movq mm2, mm5
psrlq mm5, 32
psubw mm5, mm2
movq mm3, mm7
psrlq mm7, 32
psubw mm7, mm3
punpckldq mm5, mm7

movq mm6, mm4
movq mm7, mm5
psraw mm6, 15
psraw mm7, 15
pxor mm4, mm6
pxor mm5, mm7
psubsw mm4, mm6
psubsw mm5, mm7
packuswb mm4, mm5
movq [edi], mm4

sub esi, eax
sub esi, eax
add esi, 8
add edi, 8

dec              ecx
jnz            L10402
mov esi, ebx
movd edi, mm1
add esi, eax
add edi, eax
dec              edx
jnz            L10400

emms
popa
}
#else
asm volatile
("pusha		     \n\t" "pxor      %%mm0, %%mm0 \n\t"
"mov          %3, %%eax \n\t"

"mov          %1, %%esi \n\t"
"mov          %0, %%edi \n\t"
"add       %%eax, %%edi \n\t"
"inc              %%edi \n\t"
"mov          %2, %%edx \n\t"
"sub          $2, %%edx \n\t"

".L10400:                \n\t" "mov       %%eax, %%ecx \n\t"
"shr          $3, %%ecx \n\t"
"mov       %%esi, %%ebx \n\t"
"movd      %%edi, %%mm1 \n\t"
".align 16              \n\t"
".L10402:               \n\t"

"movq    (%%esi), %%mm4 \n\t"
"movq      %%mm4, %%mm5 \n\t"
"add          $2, %%esi \n\t"
"punpcklbw %%mm0, %%mm4 \n\t"
"punpckhbw %%mm0, %%mm5 \n\t"
"movq    (%%esi), %%mm6 \n\t"
"movq      %%mm6, %%mm7 \n\t"
"sub          $2, %%esi \n\t"
"punpcklbw %%mm0, %%mm6 \n\t"
"punpckhbw %%mm0, %%mm7 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%esi), %%mm2 \n\t"
"movq      %%mm2, %%mm3 \n\t"
"add          $2, %%esi \n\t"
"punpcklbw %%mm0, %%mm2 \n\t"
"punpckhbw %%mm0, %%mm3 \n\t"
"paddw     %%mm2, %%mm4 \n\t"
"paddw     %%mm3, %%mm5 \n\t"
"paddw     %%mm2, %%mm4 \n\t"
"paddw     %%mm3, %%mm5 \n\t"
"movq    (%%esi), %%mm2 \n\t"
"movq      %%mm2, %%mm3 \n\t"
"sub          $2, %%esi \n\t"
"punpcklbw %%mm0, %%mm2 \n\t"
"punpckhbw %%mm0, %%mm3 \n\t"
"paddw     %%mm2, %%mm6 \n\t"
"paddw     %%mm3, %%mm7 \n\t"
"paddw     %%mm2, %%mm6 \n\t"
"paddw     %%mm3, %%mm7 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%esi), %%mm2 \n\t"
"movq      %%mm2, %%mm3 \n\t"
"add          $2, %%esi \n\t"
"punpcklbw %%mm0, %%mm2 \n\t"
"punpckhbw %%mm0, %%mm3 \n\t"
"paddw     %%mm2, %%mm4 \n\t"
"paddw     %%mm3, %%mm5 \n\t"
"movq    (%%esi), %%mm2 \n\t"
"movq      %%mm2, %%mm3 \n\t"
"sub          $2, %%esi \n\t"
"punpcklbw %%mm0, %%mm2 \n\t"
"punpckhbw %%mm0, %%mm3 \n\t"
"paddw     %%mm2, %%mm6 \n\t"
"paddw     %%mm3, %%mm7 \n\t"

"movq      %%mm4, %%mm2 \n\t"
"psrlq       $32, %%mm4 \n\t"
"psubw     %%mm2, %%mm4 \n\t"
"movq      %%mm6, %%mm3 \n\t"
"psrlq       $32, %%mm6 \n\t"
"psubw     %%mm3, %%mm6 \n\t"
"punpckldq %%mm6, %%mm4 \n\t"
"movq      %%mm5, %%mm2 \n\t"
"psrlq       $32, %%mm5 \n\t"
"psubw     %%mm2, %%mm5 \n\t"
"movq      %%mm7, %%mm3 \n\t"
"psrlq       $32, %%mm7 \n\t"
"psubw     %%mm3, %%mm7 \n\t"
"punpckldq %%mm7, %%mm5 \n\t"

"movq      %%mm4, %%mm6 \n\t"
"movq      %%mm5, %%mm7 \n\t"
"psraw       $15, %%mm6 \n\t"
"psraw       $15, %%mm7 \n\t"
"pxor      %%mm6, %%mm4 \n\t"
"pxor      %%mm7, %%mm5 \n\t"
"psubsw    %%mm6, %%mm4 \n\t"
"psubsw    %%mm7, %%mm5 \n\t"
"packuswb  %%mm5, %%mm4 \n\t"
"movq    %%mm4, (%%edi) \n\t"

"sub       %%eax, %%esi \n\t"
"sub       %%eax, %%esi \n\t" "add $8,          %%esi \n\t"
"add $8,          %%edi \n\t"

"dec              %%ecx \n\t"
"jnz            .L10402 \n\t"
"mov       %%ebx, %%esi \n\t"
"movd      %%mm1, %%edi \n\t"
"add       %%eax, %%esi \n\t"
"add       %%eax, %%edi \n\t"
"dec              %%edx \n\t"
"jnz            .L10400 \n\t"

"emms                   \n\t"
"popa                   \n\t":"=m" (Dest)
:"m"(Src),
"m"(rows),
"m"(columns)
);
#endif
#endif
return (0);
} else {

return (-1);
}
}


int SDL_imageFilterSobelXShiftRight(unsigned char *Src, unsigned char *Dest, int rows, int columns,
unsigned char NRightShift)
{

if ((Src == NULL) || (Dest == NULL))
return(-1);
if ((columns < 8) || (rows < 3) || (NRightShift > 7))
return (-1);

if ((SDL_imageFilterMMXdetect())) {
//#ifdef USE_MMX
#if defined(USE_MMX) && defined(i386)
#if !defined(GCC__)
__asm
{
pusha
pxor mm0, mm0
mov eax, columns
xor ebx, ebx
mov bl, NRightShift
movd mm1, ebx

mov esi, Src
mov edi, Dest
add edi, eax
inc              edi

sub rows, 2

L10410:
mov ecx, eax
shr ecx, 3
mov ebx, esi
mov edx, edi
align 16
L10412:

movq mm4, [esi]
movq mm5, mm4
add esi, 2
punpcklbw mm4, mm0
punpckhbw mm5, mm0
psrlw mm4, mm1
psrlw mm5, mm1
movq mm6, [esi]
movq mm7, mm6
sub esi, 2
punpcklbw mm6, mm0
punpckhbw mm7, mm0
psrlw mm6, mm1
psrlw mm7, mm1
add esi, eax
movq mm2, [esi]
movq mm3, mm2
add esi, 2
punpcklbw mm2, mm0
punpckhbw mm3, mm0
psrlw mm2, mm1
psrlw mm3, mm1
paddw mm4, mm2
paddw mm5, mm3
paddw mm4, mm2
paddw mm5, mm3
movq mm2, [esi]
movq mm3, mm2
sub esi, 2
punpcklbw mm2, mm0
punpckhbw mm3, mm0
psrlw mm2, mm1
psrlw mm3, mm1
paddw mm6, mm2
paddw mm7, mm3
paddw mm6, mm2
paddw mm7, mm3
add esi, eax
movq mm2, [esi]
movq mm3, mm2
add esi, 2
punpcklbw mm2, mm0
punpckhbw mm3, mm0
psrlw mm2, mm1
psrlw mm3, mm1
paddw mm4, mm2
paddw mm5, mm3
movq mm2, [esi]
movq mm3, mm2
sub esi, 2
punpcklbw mm2, mm0
punpckhbw mm3, mm0
psrlw mm2, mm1
psrlw mm3, mm1
paddw mm6, mm2
paddw mm7, mm3

movq mm2, mm4
psrlq mm4, 32
psubw mm4, mm2
movq mm3, mm6
psrlq mm6, 32
psubw mm6, mm3
punpckldq mm4, mm6
movq mm2, mm5
psrlq mm5, 32
psubw mm5, mm2
movq mm3, mm7
psrlq mm7, 32
psubw mm7, mm3
punpckldq mm5, mm7

movq mm6, mm4
movq mm7, mm5
psraw mm6, 15
psraw mm7, 15
pxor mm4, mm6
pxor mm5, mm7
psubsw mm4, mm6
psubsw mm5, mm7
packuswb mm4, mm5
movq [edi], mm4

sub esi, eax
sub esi, eax
add esi, 8
add edi, 8

dec              ecx
jnz            L10412
mov esi, ebx
mov edi, edx
add esi, eax
add edi, eax
dec rows
jnz            L10410

emms
popa
}
#else
asm volatile
("pusha		     \n\t" "pxor      %%mm0, %%mm0 \n\t"
"mov          %3, %%eax \n\t"
"xor       %%ebx, %%ebx \n\t"
"mov           %4, %%bl \n\t"
"movd      %%ebx, %%mm1 \n\t"

"mov          %1, %%esi \n\t"
"mov          %0, %%edi \n\t"
"add       %%eax, %%edi \n\t"
"inc              %%edi \n\t"

"subl            $2, %2 \n\t"

".L10410:                \n\t" "mov       %%eax, %%ecx \n\t"
"shr          $3, %%ecx \n\t"
"mov       %%esi, %%ebx \n\t"
"mov       %%edi, %%edx \n\t"
".align 16              \n\t"
".L10412:               \n\t"

"movq    (%%esi), %%mm4 \n\t"
"movq      %%mm4, %%mm5 \n\t"
"add          $2, %%esi \n\t"
"punpcklbw %%mm0, %%mm4 \n\t"
"punpckhbw %%mm0, %%mm5 \n\t"
"psrlw     %%mm1, %%mm4 \n\t"
"psrlw     %%mm1, %%mm5 \n\t"
"movq    (%%esi), %%mm6 \n\t"
"movq      %%mm6, %%mm7 \n\t"
"sub          $2, %%esi \n\t"
"punpcklbw %%mm0, %%mm6 \n\t"
"punpckhbw %%mm0, %%mm7 \n\t"
"psrlw     %%mm1, %%mm6 \n\t"
"psrlw     %%mm1, %%mm7 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%esi), %%mm2 \n\t"
"movq      %%mm2, %%mm3 \n\t"
"add          $2, %%esi \n\t"
"punpcklbw %%mm0, %%mm2 \n\t"
"punpckhbw %%mm0, %%mm3 \n\t"
"psrlw     %%mm1, %%mm2 \n\t"
"psrlw     %%mm1, %%mm3 \n\t"
"paddw     %%mm2, %%mm4 \n\t"
"paddw     %%mm3, %%mm5 \n\t"
"paddw     %%mm2, %%mm4 \n\t"
"paddw     %%mm3, %%mm5 \n\t"
"movq    (%%esi), %%mm2 \n\t"
"movq      %%mm2, %%mm3 \n\t"
"sub          $2, %%esi \n\t"
"punpcklbw %%mm0, %%mm2 \n\t"
"punpckhbw %%mm0, %%mm3 \n\t"
"psrlw     %%mm1, %%mm2 \n\t"
"psrlw     %%mm1, %%mm3 \n\t"
"paddw     %%mm2, %%mm6 \n\t"
"paddw     %%mm3, %%mm7 \n\t"
"paddw     %%mm2, %%mm6 \n\t"
"paddw     %%mm3, %%mm7 \n\t"
"add       %%eax, %%esi \n\t"
"movq    (%%esi), %%mm2 \n\t"
"movq      %%mm2, %%mm3 \n\t"
"add          $2, %%esi \n\t"
"punpcklbw %%mm0, %%mm2 \n\t"
"punpckhbw %%mm0, %%mm3 \n\t"
"psrlw     %%mm1, %%mm2 \n\t"
"psrlw     %%mm1, %%mm3 \n\t"
"paddw     %%mm2, %%mm4 \n\t"
"paddw     %%mm3, %%mm5 \n\t"
"movq    (%%esi), %%mm2 \n\t"
"movq      %%mm2, %%mm3 \n\t"
"sub          $2, %%esi \n\t"
"punpcklbw %%mm0, %%mm2 \n\t"
"punpckhbw %%mm0, %%mm3 \n\t"
"psrlw     %%mm1, %%mm2 \n\t"
"psrlw     %%mm1, %%mm3 \n\t"
"paddw     %%mm2, %%mm6 \n\t"
"paddw     %%mm3, %%mm7 \n\t"

"movq      %%mm4, %%mm2 \n\t"
"psrlq       $32, %%mm4 \n\t"
"psubw     %%mm2, %%mm4 \n\t"
"movq      %%mm6, %%mm3 \n\t"
"psrlq       $32, %%mm6 \n\t"
"psubw     %%mm3, %%mm6 \n\t"
"punpckldq %%mm6, %%mm4 \n\t"
"movq      %%mm5, %%mm2 \n\t"
"psrlq       $32, %%mm5 \n\t"
"psubw     %%mm2, %%mm5 \n\t"
"movq      %%mm7, %%mm3 \n\t"
"psrlq       $32, %%mm7 \n\t"
"psubw     %%mm3, %%mm7 \n\t"
"punpckldq %%mm7, %%mm5 \n\t"

"movq      %%mm4, %%mm6 \n\t"
"movq      %%mm5, %%mm7 \n\t"
"psraw       $15, %%mm6 \n\t"
"psraw       $15, %%mm7 \n\t"
"pxor      %%mm6, %%mm4 \n\t"
"pxor      %%mm7, %%mm5 \n\t"
"psubsw    %%mm6, %%mm4 \n\t"
"psubsw    %%mm7, %%mm5 \n\t"
"packuswb  %%mm5, %%mm4 \n\t"
"movq    %%mm4, (%%edi) \n\t"

"sub       %%eax, %%esi \n\t"
"sub       %%eax, %%esi \n\t" "add $8,          %%esi \n\t"
"add $8,          %%edi \n\t"

"dec              %%ecx \n\t"
"jnz            .L10412 \n\t"
"mov       %%ebx, %%esi \n\t"
"mov       %%edx, %%edi \n\t"
"add       %%eax, %%esi \n\t"
"add       %%eax, %%edi \n\t"
"decl                %2 \n\t"
"jnz            .L10410 \n\t"

"emms                   \n\t"
"popa                   \n\t":"=m" (Dest)
:"m"(Src),
"m"(rows),
"m"(columns),
"m"(NRightShift)
);
#endif
#endif
return (0);
} else {

return (-1);
}
}


void SDL_imageFilterAlignStack(void)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
mov ebx, esp
sub ebx, 4
and ebx, -32
mov [ebx], esp
mov esp, ebx
}
#else
asm volatile
(
"mov       %%esp, %%ebx \n\t"
"sub          $4, %%ebx \n\t"
"and        $-32, %%ebx \n\t"
"mov     %%esp, (%%ebx) \n\t"
"mov       %%ebx, %%esp \n\t"
::);
#endif
#endif
}


void SDL_imageFilterRestoreStack(void)
{
#ifdef USE_MMX
#if !defined(GCC__)
__asm
{
mov ebx, [esp]
mov esp, ebx
}
#else
asm volatile
(
"mov     (%%esp), %%ebx \n\t"
"mov       %%ebx, %%esp \n\t"
::);
#endif
#endif
}
*/
import (
	"C"
)
