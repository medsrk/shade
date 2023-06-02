#version 330

// Output fragment color
out vec4 fragColor;

// NOTE: Add here your custom variables
uniform vec2 iResolution;
uniform float iTime;
uniform vec3 pa;
uniform vec3 pb;
uniform vec3 pc;
uniform vec3 pd;

vec3 palette(float t, vec3 a, vec3 b, vec3 c, vec3 d) {
    return a + b*cos(6.28318*(c*t+d));
}

void main() {
    vec2 uv = (gl_FragCoord.xy * 2.0 - iResolution.xy)/ iResolution.y;
    vec2 uv0 = uv;
    vec3 finalColor = vec3(0.0);


    for (float i = 0.0; i < 4.0; i++) {
       uv = fract(uv * 1.5) - 0.5;

       float d = length(uv) * exp(-length(uv0));

       float d0 = length(uv0);

       vec3 col = palette(d0 + i*.4 + iTime*.4, pa, pb, pc, pd);

       d = sin(d*8. + iTime)/8.;
       d = abs(d);

       d = 0.01 / d;

       finalColor += col * d;
   }

    fragColor = vec4(finalColor, 1.0);
}
