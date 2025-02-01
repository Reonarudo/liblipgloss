#ifndef LIPGLOSS_TYPES_H
#define LIPGLOSS_TYPES_H

#include <stdint.h>
#include <stdbool.h>

// Border structure for style borders
typedef struct {
    char* Top;
    char* Bottom;
    char* Left;
    char* Right;
    char* TopLeft;
    char* TopRight;
    char* BottomLeft;
    char* BottomRight;
    char* MiddleLeft;
    char* MiddleRight;
    char* Middle;
    char* MiddleTop;
    char* MiddleBottom;
} CBorder;

// Color types for various color formats
typedef struct {
    uint32_t r;
    uint32_t g;
    uint32_t b;
    uint32_t a;
} CRGBA;

typedef struct {
    char* Light;
    char* Dark;
} CAdaptiveColor;

typedef struct {
    char* TrueColor;
    char* ANSI256;
    char* ANSI;
} CCompleteColor;

typedef struct {
    CCompleteColor Light;
    CCompleteColor Dark;
} CCompleteAdaptiveColor;

// Style properties structure
typedef struct {
    bool Bold;
    bool Italic;
    bool Underline;
    bool Strikethrough;
    bool Reverse;
    bool Blink;
    bool Faint;
    bool ColorWhitespace;
    bool Inline;
    int Width;
    int Height;
    int MaxWidth;
    int MaxHeight;
    int TabWidth;
    int PaddingTop;
    int PaddingRight;
    int PaddingBottom;
    int PaddingLeft;
    int MarginTop;
    int MarginRight;
    int MarginBottom;
    int MarginLeft;
    double AlignHorizontal;
    double AlignVertical;
    char* Foreground;
    char* Background;
    char* MarginBackground;
    char* BorderForeground;
    char* BorderBackground;
    CBorder Border;
} CStyle;

// Position constants
// These match the float64 Position values in Go
#define POS_TOP 0.0
#define POS_BOTTOM 1.0
#define POS_CENTER 0.5
#define POS_LEFT 0.0
#define POS_RIGHT 1.0

// Color profile types
typedef enum {
    PROFILE_ASCII,
    PROFILE_ANSI,
    PROFILE_ANSI256,
    PROFILE_TRUECOLOR
} CColorProfile;

// Log levels
typedef enum {
    LOG_ERROR = 0,
    LOG_WARN = 1,
    LOG_INFO = 2,
    LOG_DEBUG = 3
} CLogLevel;

// Renderer context
typedef struct {
    void* Output;
    CColorProfile ColorProfile;
    bool HasDarkBackground;
} CRenderer;

#endif /* LIPGLOSS_TYPES_H */