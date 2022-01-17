package radarr

import (
	"time"

	"code.cloudfoundry.org/bytefmt"
	"github.com/caarlos0/env/v6"
)

type Date struct {
	time.Time
}

// Sample value for Radarr dates: "04/10/2008 00:00:00"
func (t *Date) UnmarshalText(text []byte) error {
	tt, err := time.Parse("01/02/2006 00:00:00", string(text))
	t.Time = tt
	return err
}

type Movie struct {
	ID                  string `env:"movie_id" json:"id"`
	ImdbID              string `env:"movie_imdbid" json:"imdb_id"`
	PhysicalReleaseDate Date   `env:"movie_physical_release_date" json:"physical_release_date"`
	Title               string `env:"movie_title" json:"title"`
	TmdbID              string `env:"movie_tmdbid" json:"tmdb_id"`
	Year                string `env:"movie_year" json:"year"`
	InCinemasDate       Date   `env:"movie_in_cinemas_date" json:"in_cinemas_date"`
	Path                string `env:"movie_path" json:"path"`
}

type MovieFile struct {
	ID                    string   `env:"moviefile_id" json:"id"`
	RelativePath          string   `env:"moviefile_relativepath" json:"relative_path"`
	Path                  string   `env:"moviefile_path" json:"path"`
	Quality               string   `env:"moviefile_quality" json:"quality"`
	QualityVersion        string   `env:"moviefile_qualityversion" json:"quality_version"`
	ReleaseGroup          string   `env:"moviefile_releasegroup" json:"release_group"`
	SceneName             string   `env:"moviefile_scenename" json:"scene_name"`
	SourcePath            string   `env:"moviefile_sourcepath" json:"source_path"`
	SourceFolder          string   `env:"moviefile_sourcefolder" json:"source_folder"`
	IDs                   []string `env:"moviefile_ids" envSeparator:"," json:"ids"`
	RelativePaths         []string `env:"moviefile_relativepaths" envSeparator:"|" json:"relative_paths"`
	Paths                 []string `env:"moviefile_paths" envSeparator:"|" json:"paths"`
	PreviousRelativePaths []string `env:"moviefile_previousrelativepaths" envSeparator:"|" json:"previous_relative_paths"`
	PreviousPaths         []string `env:"moviefile_previouspaths" envSeparator:"|" json:"previous_paths"`
}

type Download struct {
	Client string `env:"download_client" json:"client"`
	ID     string `env:"download_id" json:"id"`
}

type HealthIssue struct {
	Level   string `env:"health_issue_level" json:"level"`
	Message string `env:"health_issue_message" json:"message"`
	Type    string `env:"health_issue_type" json:"type"`
	Wiki    string `env:"health_issue_wiki" json:"wiki"`
}

type Release struct {
	Indexer        string `env:"release_indexer" json:"indexer"`
	Quality        string `env:"release_quality" json:"quality"`
	QualityVersion string `env:"release_qualityversion" json:"quality_version"`
	Group          string `env:"release_releasegroup" json:"group"`
	Size           uint64 `env:"release_size" json:"size"`
	FormattedSize  string `json:"formatted_size"`
	Title          string `env:"release_title" json:"title"`
}

type Update struct {
	Message         string `env:"update_message" json:"message"`
	NewVersion      string `env:"update_newversion" json:"new_version"`
	PreviousVersion string `env:"update_previousversion" json:"previous_version"`
}

// Values for this struct are populated from environment variables
// that Radarr passes to a Custom Script
//  radarr_movie_id=12321
type Event struct {
	Type                 string      `env:"eventtype" json:"type"`
	FormattedType        string      `json:"formatted_type"`
	Download             Download    `json:"download"`
	Movie                Movie       `json:"movie"`
	MovieFile            MovieFile   `json:"movie_file"`
	Release              Release     `json:"release"`
	HealthIssue          HealthIssue `json:"health_issue"`
	IsUpgrade            bool        `env:"isupgrade" json:"is_upgrade"`
	DeletedRelativePaths []string    `env:"deletedrelativepaths" envSeparator:"|" json:"deleted_relative_paths"`
	DeletedPaths         []string    `env:"deletedpaths" envSeparator:"|" json:"deleted_paths"`
}

func New() (*Event, error) {
	r := new(Event)

	if err := env.Parse(r, env.Options{Prefix: "radarr_"}); err != nil {
		return nil, err
	}

	if r.DeletedPaths == nil {
		r.DeletedPaths = make([]string, 0)
	}

	if r.DeletedRelativePaths == nil {
		r.DeletedRelativePaths = make([]string, 0)
	}

	if r.MovieFile.IDs == nil {
		r.MovieFile.IDs = make([]string, 0)
	}

	if r.MovieFile.Paths == nil {
		r.MovieFile.Paths = make([]string, 0)
	}

	if r.MovieFile.RelativePaths == nil {
		r.MovieFile.RelativePaths = make([]string, 0)
	}

	if r.MovieFile.PreviousPaths == nil {
		r.MovieFile.PreviousPaths = make([]string, 0)
	}

	if r.MovieFile.PreviousRelativePaths == nil {
		r.MovieFile.PreviousRelativePaths = make([]string, 0)
	}

	if r.Release.Size != 0 {
		r.Release.FormattedSize = bytefmt.ByteSize(r.Release.Size)
	}

	switch r.Type {
	case "Grab":
		r.FormattedType = "grabbed"
	case "Download":
		if r.IsUpgrade {
			r.FormattedType = "upgraded"
		} else {
			r.FormattedType = "downloaded"
		}
	case "MovieFileDelete":
		r.FormattedType = "deleted"
	default:
		r.FormattedType = "Test"
	}

	return r, nil
}
