package handler

import (
  "encoding/json"
  "net/http"

  "github.com/api/app/model"
  "github.com/gorilla/mux"
  "github.com/jinzhu/gorm"
)

func GetAllFamousPersons(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  famousPersons := []model.FamousPerson{}
  db.Find(&famousPersons)
  respondJSON(w, http.StatusOK, famousPersons)
}

func CreateFamousPerson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  famousPerson := model.FamousPerson{}

  decoder := json.NewDecoder(r.Body)
  if err := decoder.Decode(&famousPerson); err != nil {
    respondError(w, http.StatusBadRequest, err.Error())
    return
  }
  defer r.Body.Close()

  if err := db.Save(&famousPerson).Error; err != nil {
    respondError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondJSON(w, http.StatusCreated, famousPerson)
}

func GetFamousPerson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  name := vars["name"]
  famousPerson := getFamousPersonOr404(db, name, w, r)
  if famousPerson == nil {
    return
  }
  respondJSON(w, http.StatusOK, famousPerson)
}

func UpdateFamousPerson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  name := vars["name"]
  famousPerson := getFamousPersonOr404(db, name, w, r)
  if famousPerson == nil {
    return
  }

  decoder := json.NewDecoder(r.Body)
  if err := decoder.Decode(&famousPerson); err != nil {
    respondError(w, http.StatusBadRequest, err.Error())
    return
  }
  defer r.Body.Close()

  if err := db.Save(&famousPerson).Error; err != nil {
    respondError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondJSON(w, http.StatusOK, famousPerson)
}

func DeleteFamousPerson(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  name := vars["name"]
  famousPerson := getFamousPersonOr404(db, name, w, r)
  if famousPerson == nil {
    return
  }
  if err := db.Delete(&famousPerson).Error; err != nil {
    respondError(w, http.StatusInternalServerError, err.Error())
    return
  }
  respondJSON(w, http.StatusNoContent, nil)
}

// getFamousPersonOr404 gets a famousPerson instance if exists, or respond the 404 error otherwise
func getFamousPersonOr404(db *gorm.DB, name string, w http.ResponseWriter, r *http.Request) *model.FamousPerson {
  famousPerson := model.FamousPerson{}
  if err := db.First(&famousPerson, model.FamousPerson{Name: name}).Error; err != nil {
    respondError(w, http.StatusNotFound, err.Error())
    return nil
  }
  return &famousPerson
}
