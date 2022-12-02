/*
 *Copyright (c) 2022, kaydxh
 *
 *Permission is hereby granted, free of charge, to any person obtaining a copy
 *of this software and associated documentation files (the "Software"), to deal
 *in the Software without restriction, including without limitation the rights
 *to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *copies of the Software, and to permit persons to whom the Software is
 *furnished to do so, subject to the following conditions:
 *
 *The above copyright notice and this permission notice shall be included in all
 *copies or substantial portions of the Software.
 *
 *THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 *SOFTWARE.
 */
package options

import (
	webserver_ "github.com/kaydxh/golang/pkg/webserver"
	"github.com/kaydxh/sea/pkg/sealet/application"
	"github.com/kaydxh/sea/pkg/sealet/domain/sealet"
	local_ "github.com/kaydxh/sea/pkg/sealet/infrastructure/local"
	sealet_ "github.com/kaydxh/sea/web/app/sealet"
	"github.com/kaydxh/sea/web/modules/date"
)

func (s *CompletedServerRunOptions) installWebHandler(ws *webserver_.GenericWebServer) {
	sealetFactory, err := sealet.NewFactory(sealet.FactoryConfig{
		DateRepository: &local_.Repository{},
	})
	if err != nil {
		return
	}
	app := application.Application{
		Commands: application.Commands{
			SealetHandler: application.NewSealetHandler(sealetFactory),
		},
	}
	dateCtrl := date.NewController(app)
	sealet_.NewWebHandlers(ws, dateCtrl)
}
