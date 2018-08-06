package liblsdj

/*
	TODO: make project first
 */
const LSDJ_NO_ACTIVE_PROJECT byte = 0xFF

type Lsdj_sav_t struct {
	// The projects
	lsdj_project_t* projects[LSDJ_SAV_PROJECT_COUNT];

// Index of the project that is currently being edited
/*! Indices start at 0, a value of 0xFF means there is no active project */
unsigned char activeProject;

// The song in active working memory
lsdj_song_t* song;

//! Reserved empty memory
unsigned char reserved8120[30];
};

// Create/free saves
lsdj_sav_t* lsdj_sav_new(lsdj_error_t** error);
void lsdj_sav_free(lsdj_sav_t* sav);

// Deserialize a sav
lsdj_sav_t* lsdj_sav_read(lsdj_vio_t* vio, lsdj_error_t** error);
lsdj_sav_t* lsdj_sav_read_from_file(const char* path, lsdj_error_t** error);
lsdj_sav_t* lsdj_sav_read_from_memory(const unsigned char* data, size_t size, lsdj_error_t** error);

// Serialize a sav
void lsdj_sav_write(const lsdj_sav_t* sav, lsdj_vio_t* vio, lsdj_error_t** error);
void lsdj_sav_write_to_file(const lsdj_sav_t* sav, const char* path, lsdj_error_t** error);
void lsdj_sav_write_to_memory(const lsdj_sav_t* sav, unsigned char* data, size_t size, lsdj_error_t** error);

// Set the working memory song of a sav
// The sav takes ownership of the given song, so make sure you copy it first if need be!
void lsdj_sav_set_working_memory_song(lsdj_sav_t* sav, lsdj_song_t* song, unsigned char activeProject);

// Retrieve the working memory song from a sav
lsdj_song_t* lsdj_sav_get_working_memory_song(const lsdj_sav_t* sav);

// Change the working memory song by copying from one of the projects
void lsdj_sav_set_working_memory_song_from_project(lsdj_sav_t* sav, unsigned char index, lsdj_error_t** error);

// Change which song is referenced by the working memory song
void lsdj_sav_set_active_project(lsdj_sav_t* sav, unsigned char index);

// Retrieve the index of the project the working memory song represents
// If the working memory doesn't represent any project, this is LSDJ_NO_ACTIVE_PROJECT
unsigned char lsdj_sav_get_active_project(const lsdj_sav_t* sav);

// Create a project that contains the working memory song
lsdj_project_t* lsdj_project_new_from_working_memory_song(const lsdj_sav_t* sav, lsdj_error_t** error);

// Retrieve the amount of projects in the sav (should always be 32)
unsigned int lsdj_sav_get_project_count(const lsdj_sav_t* sav);

// Change one of the projects in the sav
// The sav takes ownership of the given project, so make sure you copy it first if need be!
void lsdj_sav_set_project(lsdj_sav_t* sav, unsigned char index, lsdj_project_t* project, lsdj_error_t** error);

// Retrieve one of the projects
lsdj_project_t* lsdj_sav_get_project(const lsdj_sav_t* sav, unsigned char project);
