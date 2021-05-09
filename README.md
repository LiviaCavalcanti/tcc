# Trabalho de Conclusão de Curso - UFCG

Repositório com os experimentos do estudo para trabalho de conclusão de curso [Sobre oportunidades de redução de custo para o serviço AWS S3 através de alocações personalizadas](https://drive.google.com/file/d/1iMMmnZ2LmapVOjAr9-y9l-uoGGjkXb_7/view?usp=sharing). 

### Gerando Cargas de Trabalho de um Objeto
Carga de trabalho é uma presentação dos acessos que um objeto recebeu ao longo de um mês. Para cada objeto, ela deve ter o formato:
```
{ "trace": [
        {
        "objSize":objectSize_x,
        "requests":{"day1":[number_GET1, number_POST1], ..., dayN:[number_GETN, number_POSTN]
        }
    ]
}
```
O arquivo `simul_tcc_py/example10k.json` é um examplo de carga de trabalho.
Para gerar amostras de teste execute `simul_tcc_py/sample_generator.py`. Nele existem várias opções para gerar diferentes tipos de amostras. Elas variam tanto no tipo da distribuição usada, quanto nos perfis abordados pelo TCC.
As opções `BY_FREQUENCY`, `BY_DISTRIBUTION`, `INFREQUENT` escrevem suas saídas no terminal. Elas servem de teste para gerar uma amostra de carga de trabalho desejada. 

Para modificar uma carga de trabalho já existente em um arquivo `.json`, use a opção `MULTIPLY`. Não se esqueça de atualizar o script com os arquivos desejados.

Em `simul_tcc_py/`, também tem análises exploratórias feitas sobre o assunto. 

### Avaliando o preço de uma carga de trabalho

Um simulador de custo do S3 está em `tree_algorithm.go`. Para executar a carga desejada, mude o arquivo de entrada definido no próprio script.
A saída mostra o custo da carga de trabalho no Infrequent Access do S3 e do modelo proposto no TCC.


Obs.: Ambos os scripts podem ser reproduzidos por
```
python simul_tcc_py/sample_generator.py
go tree_algorithm.go
```
Caso seja necessário alguma mudança, modifique o arquivo. 
