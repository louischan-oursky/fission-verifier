package verifier

import (
	"context"

	"github.com/fission/fission"
	"github.com/fission/fission/environments/fetcher"
	"go.uber.org/zap"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type FetchOptions struct {
	// The following are options for creating
	Dir       string
	SecretDir string
	ConfigDir string
	// The real arguments
	PackageNamespace string
	PackageName      string
}

func Fetch(options FetchOptions) error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}
	defer func() {
		if logger != nil {
			logger.Sync()
		}
	}()

	f, err := fetcher.MakeFetcher(logger, options.Dir, options.SecretDir, options.ConfigDir)
	if err != nil {
		return err
	}

	req := fission.FunctionFetchRequest{
		FetchType: fission.FETCH_DEPLOYMENT,
		Package: metav1.ObjectMeta{
			Namespace: options.PackageNamespace,
			Name:      options.PackageName,
		},
	}
	ctx := context.Background()
	_, err = f.Fetch(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
