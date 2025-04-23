package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatus(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/status", nil)
	w := httptest.NewRecorder()
	HandleStatus(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestClick(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/click", nil)
	w := httptest.NewRecorder()
	HandleClick(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestUpgrade(t *testing.T) {
	body := []byte(`{"type":"attack"}`)
	req := httptest.NewRequest(http.MethodPost, "/upgrade", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	HandleUpgrade(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusPaymentRequired {
		t.Errorf("Expected 200 or 402, got %d", w.Code)
	}
}

func TestBattle(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/battle", nil)
	w := httptest.NewRecorder()
	HandleBattle(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestHeal(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/heal", nil)
	w := httptest.NewRecorder()
	HandleHeal(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusPaymentRequired && w.Code != http.StatusConflict {
		t.Errorf("Expected 200/402/409, got %d", w.Code)
	}
}

func TestCollect(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/collect", nil)
	w := httptest.NewRecorder()
	HandleCollect(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestShop(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/shop", nil)
	w := httptest.NewRecorder()
	HandleShop(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestInventory(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/inventory", nil)
	w := httptest.NewRecorder()
	HandleInventory(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", w.Code)
	}
}

func TestBuy(t *testing.T) {
	// Пробуем купить Bronze Sword
	body, _ := json.Marshal(Item{Name: "Bronze Sword"})
	req := httptest.NewRequest(http.MethodPost, "/buy", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	HandleBuy(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusBadRequest {
		t.Errorf("Expected 200 or 400, got %d", w.Code)
	}
}

func TestEquip(t *testing.T) {
	body, _ := json.Marshal(EquipRequest{Name: "Bronze Sword"})
	req := httptest.NewRequest(http.MethodPost, "/equip", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	HandleEquip(w, req)

	if w.Code != http.StatusOK && w.Code != http.StatusNotFound {
		t.Errorf("Expected 200 or 404, got %d", w.Code)
	}
}
