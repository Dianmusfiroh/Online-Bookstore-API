
package middleware

import (
	"fmt"
	"os" // PERBAIKAN: Impor 'os' untuk membaca environment variable
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Struct ini sudah benar, tidak perlu diubah.
type JwtClaims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}
func JWTAuth() fiber.Handler {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		// Panic akan menghentikan server jika secret tidak disetel, ini lebih aman.
		panic("FATAL ERROR: JWT_SECRET environment variable is not set")
	}

	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing authorization header"})
		}

		// Memotong prefix "Bearer " sudah benar.
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
		}
		tokenStr := parts[1]

		// PERBAIKAN: Buat instance dari struct claims Anda untuk diisi.
		claims := &JwtClaims{}

		// PERBAIKAN: Gunakan jwt.ParseWithClaims untuk mem-parsing token ke dalam struct Anda.
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			// PERBAIKAN: Tambahkan validasi metode signing untuk keamanan ekstra.
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired token"})
		}
		
		// Baris ini untuk debugging, untuk memastikan claims sudah terisi dengan benar.
		fmt.Printf("[Middleware] Claims parsed successfully: UserID=%d, Role=%s\n", claims.UserID, claims.Role)

		// PERBAIKAN: Simpan object claims yang SUDAH TERISI ke context.
		c.Locals("claims", claims)

		return c.Next()
	}
}

func AdminAuth() fiber.Handler {
    return func(c *fiber.Ctx) error {
        claims, ok := c.Locals("claims").(*JwtClaims)
        if !ok {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve user claims"})
        }
        
        if claims.Role != "admin" {
            return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"error": "Forbidden: Admin access required"})
        }
        return c.Next()
    }
}
