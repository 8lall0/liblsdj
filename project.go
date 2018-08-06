package liblsdj

/*
	TODO: make song first
 */
// The length of project names
const (
	LSDJ_PROJECT_NAME_LENGTH int = 8
	BLOCK_SIZE byte = 0x200
	BLOCK_COUNT byte = 191
)

// Representation of a project within an LSDJ sav file
type Lsdj_project_t struct {
	// The name of the project
	name [LSDJ_PROJECT_NAME_LENGTH]string
	// The version of the project
	version [5]string;
	// The song belonging to this project
	/*! If this is NULL, the project isn't in use */
	song *lsdj_song_t;
};

// Create/free projects
lsdj_project_t* lsdj_project_new(lsdj_error_t** error);
void lsdj_project_free(lsdj_project_t* project);

// Deserialize a project from LSDSNG
lsdj_project_t* lsdj_project_read_lsdsng(lsdj_vio_t* vio, lsdj_error_t** error);
lsdj_project_t* lsdj_project_read_lsdsng_from_file(const char* path, lsdj_error_t** error);
lsdj_project_t* lsdj_project_read_lsdsng_from_memory(const unsigned char* data, size_t size, lsdj_error_t** error);

// Write a project to an lsdsng file
void lsdj_project_write_lsdsng(const lsdj_project_t* project, lsdj_vio_t* vio, lsdj_error_t** error);
void lsdj_project_write_lsdsng_to_file(const lsdj_project_t* project, const char* path, lsdj_error_t** error);
void lsdj_project_write_lsdsng_to_memory(const lsdj_project_t* project, unsigned char* data, size_t size, lsdj_error_t** error);

// Change data in a project
void lsdj_project_set_name(lsdj_project_t* project, const char* data, size_t size);
void lsdj_project_get_name(const lsdj_project_t* project, char* data, size_t size);
void lsdj_project_set_version(lsdj_project_t* project, unsigned char version);
unsigned char lsdj_project_get_version(const lsdj_project_t* project);
void lsdj_project_set_song(lsdj_project_t* project, lsdj_song_t* song);
lsdj_song_t* lsdj_project_get_song(const lsdj_project_t* project);