Pod::Spec.new do |spec|
  spec.name         = 'Energi Core'
  spec.version      = '{{.Version}}'
  spec.license      = { :type => 'GNU Lesser General Public License, Version 3.0' }
  spec.homepage     = 'https://github.com/ethereum/go-ethereum'
  spec.authors      = { {{range .Contributors}}
		'{{.Name}}' => '{{.Email}}',{{end}}
	}
  spec.summary      = 'iOS Energi Client'
  spec.source       = { :git => 'https://github.com/ethereum/go-ethereum.git', :commit => '{{.Commit}}' }

	spec.platform = :ios
  spec.ios.deployment_target  = '9.0'
	spec.ios.vendored_frameworks = 'Frameworks/EnergiCore.framework'

	spec.prepare_command = <<-CMD
    curl https://energi3store.blob.core.windows.net/builds/{{.Archive}}.tar.gz | tar -xvz
    mkdir Frameworks
    mv {{.Archive}}/EnergiCore.framework Frameworks
    rm -rf {{.Archive}}
  CMD
end
